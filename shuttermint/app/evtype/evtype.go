// evtype declares the the different event types sent by shuttermint
package evtype

var (
	CheckIn             = "shutter.check-in"
	PolyCommitment      = "shutter.poly-commitment-registered"
	PolyEval            = "shutter.poly-eval-registered"
	BatchConfig         = "shutter.batch-config"
	DecryptionSignature = "shutter.decryption-signature"
	EonStarted          = "shutter.eon-started"
	Accusation          = "shutter.accusation-registered"
	Apology             = "shutter.apology-registered"
)
