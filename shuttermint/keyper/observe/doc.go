// Package observe contains the MainChain and Shutter structs, which the keyper uses to fetch the
// necessary information from the ethereum node and the shuttermint node. The SyncToHead methods
// can be used to fetch the latest information. All other public methods do not modify the stored
// data. Do not mutate any of the data stored in these structs.
package observe
