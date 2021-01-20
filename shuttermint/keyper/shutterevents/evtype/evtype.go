// evtype declares the the different event types sent by shuttermint
package evtype

var (
	Accusation          = "shutter.accusation-registered"
	Apology             = "shutter.apology-registered"
	BatchConfig         = "shutter.batch-config"
	CheckIn             = "shutter.check-in"
	DecryptionSignature = "shutter.decryption-signature"
	EonStarted          = "shutter.eon-started"
	PolyCommitment      = "shutter.poly-commitment-registered"
	PolyEval            = "shutter.poly-eval-registered"
	EpochSecretKeyShare = "shutter.epoch-secret-key-share"
)
