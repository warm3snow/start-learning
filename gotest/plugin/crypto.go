package plugin1

type IHSMAdapter interface {
	// for PKCS11
	GetSM2KeyId(keyIdex int, isPrivate bool) (string, error)
	GetRSAKeyId(keyIdex int, isPrivate bool) (string, error)
	GetECCKeyId(keyIdex int, isPrivate bool) (string, error)

	GetSM4KeyIdex(keyIdex int) (string, error)
	GetAESKeyId(keyIdex int) (string, error)

	// GetSM3SM2CKM this is a exceptional interface
	GetSM3SM2CKM() uint

	// For SDF
	GetSM2KeyAccessRight(keyIdex int) (newKeyIdex int, need bool)
	GetSM4KeyAccessRight(keyIdex int) (newKeyIdex int, need bool)

	GetRSAKeyAccessRight(keyIdex int) (newKeyIdex int, need bool)
	GetAESKeyAccessRight(keyIdex int) (newKeyIdex int, need bool)
}
