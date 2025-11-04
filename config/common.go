package config

type CommonConfig struct {
	BackgroundRsaPrivateKey string `mapstructure:"background_rsa_private_key" json:"background_rsa_private_key" yaml:"background_rsa_private_key"` // 后端解密RSA私钥
	BackgroundRsaPublicKey  string `mapstructure:"background_rsa_public_key" json:"background_rsa_public_key" yaml:"background_rsa_public_key"`    // 后端加密RSA公钥
	FrontRsaPrivateKey      string `mapstructure:"front_rsa_private_key" json:"front_rsa_private_key" yaml:"front_rsa_private_key"`                // 前端解密RSA私钥
	FrontRsaPublicKey       string `mapstructure:"front_rsa_public_key" json:"front_rsa_public_key" yaml:"front_rsa_public_key"`                   // 前端加密RSA公钥
	SqlMd5CryptKey          string `mapstructure:"sql_md5_crypt_key" json:"sql_md5_crypt_key" yaml:"sql_md5_crypt_key"`                            // 数据库MD5加密密钥
	AesCryptKey             string `mapstructure:"aes_crypt_key" json:"aes_crypt_key" yaml:"aes_crypt_key"`                                        // 数据库中AES算法密钥
	TokenCryptKey           string `mapstructure:"token_crypt_key" json:"token_crypt_key" yaml:"token_crypt_key"`                                  // jwt加密密钥
	TokenSignKey            string `mapstructure:"token_sign_key" json:"token_sign_key" yaml:"token_sign_key"`                                     // jwt内sign加密密钥
	CommonHashKey           string `mapstructure:"common_hash_key" json:"common_hash_key" yaml:"common_hash_key"`                                  // 常规hash加密密钥
	PublishingPayChannelId  int    `mapstructure:"publishing_pay_channel_id" json:"publishing_pay_channel_id" yaml:"publishing_pay_channel_id"`    // 发行默认充值渠道ID
	GameHashKey             string `mapstructure:"game_hash_key" json:"game_hash_key" yaml:"game_hash_key"`                                        // 游戏加密密钥
	CtxTokenDataKey         string `mapstructure:"ctx_token_data_key" json:"ctx_token_data_key" yaml:"ctx_token_data_key"`                         // 上下文中存储token的键值
	AuthorizationHeadKey    string `mapstructure:"authorization_head_key" json:"authorization_head_key" yaml:"authorization_head_key"`             // 授权头部键值
	CtxRequestIdKey         string `mapstructure:"ctx_request_id_key" json:"ctx_request_id_key" yaml:"ctx_request_id_key"`                         // 追踪链中的上下文键值
}
