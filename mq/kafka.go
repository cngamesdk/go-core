package mq

import (
	"context"
	"encoding/json"
	"github.com/duke-git/lancet/v2/random"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/scram"
	"go.uber.org/zap"
	"time"
)

type KafkaConfig struct {
	Brokers     []string
	TopicPrefix string
	GroupID     string
	Username    string
	Password    string
	Mechanism   string
}

type KafkaClient struct {
	config    KafkaConfig
	logger    *zap.Logger
	producer  *kafka.Writer
	consumers map[string]*kafka.Reader
}

type Message struct {
	ID        string                 `json:"id"`
	Topic     string                 `json:"topic"`
	Action    string                 `json:"action"`
	Payload   map[string]interface{} `json:"payload"`
	Timestamp time.Time              `json:"timestamp"`
	Retry     int                    `json:"retry"`
}

func (c *KafkaClient) Publish(ctx context.Context, topic, action string, payload interface{}) error {
	uuid, _ := random.UUIdV4()
	message := Message{
		ID:        uuid,
		Topic:     topic,
		Action:    action,
		Timestamp: time.Now(),
		Retry:     0,
	}

	// 转换payload
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		var payloadMap map[string]interface{}
		if err := json.Unmarshal(data, &payloadMap); err != nil {
			return err
		}
		message.Payload = payloadMap
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// 发送到Kafka
	err = c.producer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(message.ID),
		Value: data,
		Topic: c.config.TopicPrefix + "." + topic,
		Time:  time.Now(),
	})

	if err != nil {
		c.logger.Error("Failed to publish message",
			zap.String("topic", topic),
			zap.String("action", action),
			zap.Error(err))
		return err
	}

	c.logger.Info("Message published",
		zap.String("topic", topic),
		zap.String("action", action),
		zap.String("message_id", message.ID))

	return nil
}

func (c *KafkaClient) Subscribe(ctx context.Context, topic string, handler func(Message) error) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  c.config.Brokers,
		Topic:    c.config.TopicPrefix + "." + topic,
		GroupID:  c.config.GroupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  time.Second,
	})

	c.consumers[topic] = reader

	go func() {
		for {
			select {
			case <-ctx.Done():
				_ = reader.Close()
				return
			default:
				m, err := reader.ReadMessage(ctx)
				if err != nil {
					c.logger.Error("Failed to read message",
						zap.String("topic", topic),
						zap.Error(err))
					continue
				}

				var message Message
				if err := json.Unmarshal(m.Value, &message); err != nil {
					c.logger.Error("Failed to unmarshal message",
						zap.String("topic", topic),
						zap.Error(err))
					continue
				}

				// 处理消息
				if err := handler(message); err != nil {
					c.logger.Error("Failed to handle message",
						zap.String("topic", topic),
						zap.String("message_id", message.ID),
						zap.Error(err))

					// 重试逻辑
					if message.Retry < 3 {
						message.Retry++
						c.retryMessage(ctx, message)
					}
				}
			}
		}
	}()

	return nil
}

func (c *KafkaClient) retryMessage(ctx context.Context, message Message) {
	// 延迟重试
	time.Sleep(time.Duration(message.Retry) * time.Second)

	data, _ := json.Marshal(message)
	_ = c.producer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(message.ID),
		Value: data,
		Topic: c.config.TopicPrefix + "." + message.Topic,
		Time:  time.Now(),
	})
}

func NewClient(config KafkaConfig, logger *zap.Logger) (*KafkaClient, error) {
	// 配置SASL认证
	var mechanism sasl.Mechanism
	if config.Username != "" && config.Password != "" {
		m, err := scram.Mechanism(scram.SHA512, config.Username, config.Password)
		if err != nil {
			return nil, err
		}
		mechanism = m
	}

	// 创建生产者
	producer := &kafka.Writer{
		Addr:         kafka.TCP(config.Brokers...),
		Topic:        config.TopicPrefix + ".commands",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		Transport: &kafka.Transport{
			SASL: mechanism,
		},
	}

	return &KafkaClient{
		config:    config,
		logger:    logger,
		producer:  producer,
		consumers: make(map[string]*kafka.Reader),
	}, nil
}
