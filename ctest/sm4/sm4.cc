#include <openssl/sms4.h>

int main(void) {
    //加密参数初始化
    sms4_key_t sms4_key_enc;
    unsigned char *plain_text = { 0 };
    char *key = "01234567891234560123456789123456";
    unsigned char *iv = "0123456789123456";
    //SM4加密
    memcpy(sms4_key.rk, key, 32);
    sms4_set_encrypt_key(&sms4_key, iv);
    sms4_cbc_encrypt(plain_text, enc_text, 64, sms4_key_enc.rk, iv, 1);

    //解密参数初始化
    sms4_key_t sms4_key_decrypt;
    //SM4解密
    memcpy(sms4_key.rk, key, 32);
    sms4_set_decrypt_key(reinterpret_cast<sms4_key_t *>(sms4_key_decrypt->rk), iv);
    sms4_cbc_encrypt((uint8_t *) data, plaintext, 64,
            reinterpret_cast<const sms4_key_t *>(sms4_key_decrypt->rk), iv, 0);
    return 0;
}