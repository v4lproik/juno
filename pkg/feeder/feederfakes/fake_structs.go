package feederfakes

// notest
import (
	"encoding/json"

	feeder "github.com/NethermindEth/juno/pkg/feeder"
	abi "github.com/NethermindEth/juno/pkg/feeder/abi"
)

//Returns feeder_gateway response for contract_addresses
func ReturnFakeContractAddressInfo() *feeder.ContractAddresses {
	var contractAddresses *feeder.ContractAddresses
	rjson := "{\"Starknet\": \"0xc662c410C0ECf747543f5bA90660f6ABeBD9C8c4\", \"GpsStatementVerifier\": \"0x47312450B3Ac8b5b8e247a6bB6d523e7605bDb60\"}"
	bjson := []byte(rjson)
	json.Unmarshal(bjson, &contractAddresses)
	return contractAddresses
}

//Returns feeder_gateway response for state_update on blockNumber = 2
//blockHash=0x4e1f77f39545afe866ac151ac908bd1a347a2a8a7d58bef1276db4f06fdf2f6
func ReturnFakeStateUpdateInfo() *feeder.StateUpdateResponse {
	var stateUpdate *feeder.StateUpdateResponse
	rjson := "{\"block_hash\": \"0x4e1f77f39545afe866ac151ac908bd1a347a2a8a7d58bef1276db4f06fdf2f6\", \"new_root\": \"03ceee867d50b5926bb88c0ec7e0b9c20ae6b537e74aac44b8fcf6bb6da138d9\", \"old_root\": \"0525aed4da9cc6cce2de31ba79059546b0828903279e4eaa38768de33e2cac32\", \"state_diff\": {\"storage_diffs\": {\"0x1fb4457f3fe8a976bdb9c04dd21549beeeb87d3867b10effe0c4bd4064a8e4\": [{\"key\": \"0x56c060e7902b3d4ec5a327f1c6e083497e586937db00af37fe803025955678f\", \"value\": \"0x75495b43f53bd4b9c9179db113626af7b335be5744d68c6552e3d36a16a747c\"}], \"0x5790719f16afe1450b67a92461db7d0e36298d6a5f8bab4f7fd282050e02f4f\": [{\"key\": \"0x772c29fae85f8321bb38c9c3f6edb0957379abedc75c17f32bcef4e9657911a\", \"value\": \"0x6d4ca0f72b553f5338a95625782a939a49b98f82f449c20f49b42ec60ed891c\"}], \"0x57b973bf2eb26ebb28af5d6184b4a044b24a8dcbf724feb95782c4d1aef1ca9\": [{\"key\": \"0x4f2c206f3f2f1380beeb9fe4302900701e1cb48b9b33cbe1a84a175d7ce8b50\", \"value\": \"0x2a614ae71faa2bcdacc5fd66965429c57c4520e38ebc6344f7cf2e78b21bd2f\"}], \"0x2d6c9569dea5f18628f1ef7c15978ee3093d2d3eec3b893aac08004e678ead3\": [{\"key\": \"0x7f93985c1baa5bd9b2200dd2151821bd90abb87186d0be295d7d4b9bc8ca41f\", \"value\": \"0x127cd00a078199381403a33d315061123ce246c8e5f19aa7f66391a9d3bf7c6\"}]}, \"deployed_contracts\": [{\"address\": \"0x1fb4457f3fe8a976bdb9c04dd21549beeeb87d3867b10effe0c4bd4064a8e4\", \"contract_hash\": \"010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\"}, {\"address\": \"0x5790719f16afe1450b67a92461db7d0e36298d6a5f8bab4f7fd282050e02f4f\", \"contract_hash\": \"010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\"}, {\"address\": \"0x57b973bf2eb26ebb28af5d6184b4a044b24a8dcbf724feb95782c4d1aef1ca9\", \"contract_hash\": \"010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\"}, {\"address\": \"0x2d6c9569dea5f18628f1ef7c15978ee3093d2d3eec3b893aac08004e678ead3\", \"contract_hash\": \"010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\"}]}}"
	bjson := []byte(rjson)
	json.Unmarshal(bjson, &stateUpdate)
	return stateUpdate
}

////Returns feeder_gateway response for get_code on
//contractAddress=0x0090bff87efa37c2a8d0fd8b903ca1220dbb375ad18346e0da59af7f7e6c4285
func ReturnFakeCodeInfo() *feeder.CodeInfo {
	var FakeCodeInfo feeder.CodeInfo
	var FakeAbi abi.Abi
	FakeAbi = *ReturnAbiInfo()
	FakeCodeInfo.Abi = FakeAbi
	FakeCodeInfo.Bytecode = []string{"0x20780017fff7ffd", "0x4", "0x400780017fff7ffd", "0x1", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x44656c656761746543616c6c", "0x400280007ff97fff", "0x400380017ff97ffa", "0x400380027ff97ffb", "0x400380037ff97ffc", "0x400380047ff97ffd", "0x482680017ff98000", "0x7", "0x480280057ff98000", "0x480280067ff98000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x44656c65676174654c3148616e646c6572", "0x400280007ff97fff", "0x400380017ff97ffa", "0x400380027ff97ffb", "0x400380037ff97ffc", "0x400380047ff97ffd", "0x482680017ff98000", "0x7", "0x480280057ff98000", "0x480280067ff98000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x53746f7261676552656164", "0x400280007ffc7fff", "0x400380017ffc7ffd", "0x482680017ffc8000", "0x3", "0x480280027ffc8000", "0x208b7fff7fff7ffe", "0x480680017fff8000", "0x53746f726167655772697465", "0x400280007ffb7fff", "0x400380017ffb7ffc", "0x400380027ffb7ffd", "0x482680017ffb8000", "0x3", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x480680017fff8000", "0xf920571b9f85bdd92a867cfdc73319d0f8836f0e69e06e4c5566b6203f75cc", "0x208b7fff7fff7ffe", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffffa", "0x480a7ffb7fff8000", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe6", "0x48127ffe7fff8000", "0x48127ff57fff8000", "0x48127ff57fff8000", "0x48127ffc7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffed", "0x480a7ffa7fff8000", "0x48127ffe7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe0", "0x48127ff67fff8000", "0x48127ff67fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe5", "0x208b7fff7fff7ffe", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffaf", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe8", "0x208b7fff7fff7ffe", "0x480a7ffa7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff3", "0x208b7fff7fff7ffe", "0x482680017ffd8000", "0x1", "0x402a7ffd7ffc7fff", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480280007ffd8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff3", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480680017fff8000", "0x0", "0x48127ffb7fff8000", "0x208b7fff7fff7ffe", "0x480a7ff87fff8000", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffd5", "0x48127ffc7fff8000", "0x48127ffe7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff88", "0x48127ffd7fff8000", "0x48127ff17fff8000", "0x48127ff17fff8000", "0x48127ffb7fff8000", "0x48127ffb7fff8000", "0x208b7fff7fff7ffe", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480a7ffa7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffe9", "0x208b7fff7fff7ffe", "0x480a7ff87fff8000", "0x480a7ff97fff8000", "0x480a7ffa7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffba", "0x48127ffc7fff8000", "0x48127ffe7fff8000", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff79", "0x48127ffd7fff8000", "0x48127ff17fff8000", "0x48127ff17fff8000", "0x208b7fff7fff7ffe", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x480a7ffa7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffeb", "0x40780017fff7fff", "0x1", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x48127ffc7fff8000", "0x480680017fff8000", "0x0", "0x48127ffb7fff8000", "0x208b7fff7fff7ffe", "0x480a7ffb7fff8000", "0x480a7ffc7fff8000", "0x480a7ffd7fff8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffff99", "0x208b7fff7fff7ffe", "0x40780017fff7fff", "0x1", "0x4003800080007ffc", "0x4826800180008000", "0x1", "0x480a7ffd7fff8000", "0x4828800080007ffe", "0x480a80007fff8000", "0x208b7fff7fff7ffe", "0x402b7ffd7ffc7ffd", "0x480280007ffb8000", "0x480280017ffb8000", "0x480280027ffb8000", "0x1104800180018000", "0x800000000000010ffffffffffffffffffffffffffffffffffffffffffffffee", "0x48127ffe7fff8000", "0x1104800180018000", "0x800000000000010fffffffffffffffffffffffffffffffffffffffffffffff1", "0x48127ff47fff8000", "0x48127ff47fff8000", "0x48127ffb7fff8000", "0x48127ffb7fff8000", "0x48127ffb7fff8000", "0x208b7fff7fff7ffe"}

	return &FakeCodeInfo

}

//returns ABI for contractAddress=0x0090bff87efa37c2a8d0fd8b903ca1220dbb375ad18346e0da59af7f7e6c4285
func ReturnAbiInfo() *abi.Abi {
	_json := "[{\"inputs\": [{\"name\": \"implementation\", \"type\": \"felt\"}], \"name\": \"constructor\", \"outputs\": [], \"type\": \"constructor\"}, {\"inputs\": [{\"name\": \"selector\", \"type\": \"felt\"}, {\"name\": \"calldata_size\", \"type\": \"felt\"}, {\"name\": \"calldata\", \"type\": \"felt*\"}], \"name\": \"__default__\", \"outputs\": [{\"name\": \"retdata_size\", \"type\": \"felt\"}, {\"name\": \"retdata\", \"type\": \"felt*\"}], \"type\": \"function\"}, {\"inputs\": [{\"name\": \"selector\", \"type\": \"felt\"}, {\"name\": \"calldata_size\", \"type\": \"felt\"}, {\"name\": \"calldata\", \"type\": \"felt*\"}], \"name\": \"__l1_default__\", \"outputs\": [], \"type\": \"l1_handler\"}, {\"inputs\": [], \"name\": \"get_implementation\", \"outputs\": [{\"name\": \"implementation\", \"type\": \"felt\"}], \"stateMutability\": \"view\", \"type\": \"function\"}]"
	bjson := []byte(_json)
	var a abi.Abi
	a.UnmarshalAbiJSON(bjson)

	return &a
}

//returns ABI with full coverage
func ReturnAbiInfo_Full() *abi.Abi {
	_json := "[{\"inputs\": [{\"name\": \"funct\", \"type\": \"felt\"}], \"name\": \"function-custom\", \"outputs\": [], \"type\": \"function\"}, {\"inputs\": [{\"name\": \"implementation\", \"type\": \"felt\"}], \"name\": \"L1Handler-custom\", \"outputs\": [], \"type\": \"l1_handler\"}, {\"members\": [{\"offset\": 1, \"name\": \"member\", \"type\": \"struct\"}], \"name\": \"Struct-custom\", \"size\": 3, \"type\": \"struct\"}, {\"inputs\": [{\"name\": \"constr\", \"type\": \"felt\"}], \"name\": \"constructor-custom\", \"outputs\": [], \"type\": \"constructor\"}, {\"data\": [{\"name\": \"storage_cells_len\", \"type\": \"felt\"}, {\"name\": \"storage_cells\", \"type\": \"StorageCell*\"}], \"keys\": [], \"name\": \"log_storage_cells\", \"type\": \"event\"}]"
	bjson := []byte(_json)
	var a abi.Abi
	a.UnmarshalAbiJSON(bjson)

	return &a
}

//returns ABI with full coverage
func ReturnAbiInfo_Fail() error {
	_json := "[{\"inputs\": [{\"name\": \"funct\", \"type\": \"felt\"}], \"name\": \"function-custom\", \"outputs\": [], \"type\": \"function\"}, {\"inputs\": [{\"name\": \"implementation\", \"type\": \"unknown\"}], \"name\": \"L1Handler-custom\", \"outputs\": [], \"type\": \"l1_handler\"}, {\"members\": [{\"offset\": 1, \"name\": \"member\", \"type\": \"struct\"}], \"name\": \"Struct-custom\", \"size\": 3, \"type\": \"struct\"}, {\"inputs\": [{\"name\": \"constr\", \"type\": \"felt\"}], \"name\": \"constructor-custom\", \"outputs\": [], \"type\": \"constructor\"}, {\"data\": [{\"name\": \"storage_cells_len\", \"type\": \"felt\"}, {\"name\": \"storage_cells\", \"type\": \"StorageCell*\"}], \"keys\": [], \"name\": \"log_storage_cells\", \"type\": \"event\"}]"
	bjson := []byte(_json)
	var a abi.Abi
	err := a.UnmarshalAbiJSON(bjson)

	return err
}

//Block info for blockNumber=0
func ReturnFakeBlockInfo() *feeder.StarknetBlock {
	var block feeder.StarknetBlock
	rjson := "{\"block_hash\": \"0x47c3637b57c2b079b93c61539950c17e868a28f46cdef28f88521067f21e943\", \"parent_block_hash\": \"0x0\", \"block_number\": 0, \"state_root\": \"021870ba80540e7831fb21c591ee93481f5ae1bb71ff85a86ddd465be4eddee6\", \"status\": \"ACCEPTED_ON_L1\", \"gas_price\": \"0x174876e800\", \"transactions\": [{\"contract_address\": \"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"contract_address_salt\": \"0x546c86dc6e40a5e5492b782d8964e9a4274ff6ecb16d31eb09cee45a3564015\", \"class_hash\": \"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\", \"constructor_calldata\": [\"0x6cf6c2f36d36b08e591e4489e92ca882bb67b9c39a3afccf011972a8de467f0\", \"0x7ab344d88124307c07b56f6c59c12f4543e9c96398727854a322dea82c73240\"], \"transaction_hash\": \"0xe0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75\", \"type\": \"DEPLOY\"}, {\"contract_address\": \"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"contract_address_salt\": \"0x12afa0f342ece0468ca9810f0ea59f9c7204af32d1b8b0d318c4f2fe1f384e\", \"class_hash\": \"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\", \"constructor_calldata\": [\"0xcfc2e2866fd08bfb4ac73b70e0c136e326ae18fc797a2c090c8811c695577e\", \"0x5f1dd5a5aef88e0498eeca4e7b2ea0fa7110608c11531278742f0b5499af4b3\"], \"transaction_hash\": \"0x12c96ae3c050771689eb261c9bf78fac2580708c7f1f3d69a9647d8be59f1e1\", \"type\": \"DEPLOY\"}, {\"contract_address\": \"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"entry_point_selector\": \"0x12ead94ae9d3f9d2bdb6b847cf255f1f398193a1f88884a0ae8e18f24a037b6\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0xc84dd7fd43a7defb5b7a15c4fbbe11cbba6db1ba\"], \"signature\": [], \"transaction_hash\": \"0xce54bbc5647e1c1ea4276c01a708523f740db0ff5474c77734f73beec2624\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"entry_point_selector\": \"0x218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"0x0\"], \"signature\": [], \"transaction_hash\": \"0x1c924916a84ef42a3d25d29c5d1085fe212de04feadc6e88d4c7a6e5b9039bf\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"entry_point_selector\": \"0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x7dbfec95c10bbc2fd3f37a89ae6e027826134f955251d11c784a6b34fdf50\", \"0x2\", \"0x4e7e989d58a17cd279eca440c5eaa829efb6f9967aaad89022acbe644c39b36\", \"0x453ae0c9610197b18b13645c44d3d0a407083d96562e8752aab3fab616cecb0\"], \"signature\": [], \"transaction_hash\": \"0xa66c346e273cc49510ef2e1620a1a7922135cb86ab227b86e0afd12243bd90\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"entry_point_selector\": \"0x27c3334165536f239cfd400ed956eabff55fc60de4fb56728b6a4f6b87db01c\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f\", \"0x4\", \"0x4be52041fee36ab5199771acf4b5d260d223297e588654e5c9477df2efa542a\", \"0x2\", \"0x299e2f4b5a873e95e65eb03d31e532ea2cde43b498b50cd3161145db5542a5\", \"0x3d6897cf23da3bf4fd35cc7a43ccaf7c5eaf8f7c5b9031ac9b09a929204175f\"], \"signature\": [], \"transaction_hash\": \"0x5c71675616b49fb9d16cac8beaaa65f62dc5a532e92785055c15c825166dbbf\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"entry_point_selector\": \"0x218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"0x1\"], \"signature\": [], \"transaction_hash\": \"0x60e05c41a6622592a2e2eff90a9f2e495296a3be9596e7bc4dfbafce00d7a6a\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"entry_point_selector\": \"0x19a35a6e95cb7a3318dbb244f20975a1cd8587cc6b5259f15f61d7beb7ee43b\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"0x5aee31408163292105d875070f98cb48275b8c87e80380b78d30647e05854d5\"], \"signature\": [], \"transaction_hash\": \"0x5634f2847140263ba59480ad4781dacc9991d0365145489b27a198ebed2f969\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x6ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae\", \"contract_address_salt\": \"0x5098705e4d57a8620e5b387855ef4dc82f0ccd84b7299dc1179b87517249127\", \"class_hash\": \"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\", \"constructor_calldata\": [\"0x48cba68d4e86764105adcdcf641ab67b581a55a4f367203647549c8bf1feea2\", \"0x362d24a3b030998ac75e838955dfee19ec5b6eceb235b9bfbeccf51b6304d0b\"], \"transaction_hash\": \"0xb049c384cf75174150a2540835cc2abdcca1d3a3750298a1741a621983e35a\", \"type\": \"DEPLOY\"}, {\"contract_address\": \"0x735596016a37ee972c42adef6a3cf628c19bb3794369c65d2c82ba034aecf2c\", \"contract_address_salt\": \"0x60bc7461113e4af46fd52e5ecbc5c3f4be92ed7f1329d53721f9bfbc0370cc\", \"class_hash\": \"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\", \"constructor_calldata\": [\"0x2f50710449a06a9fa789b3c029a63bd0b1f722f46505828a9f815cf91b31d8\", \"0x2a222e62eabe91abdb6838fa8b267ffe81a6eb575f61e96ec9aa4460c0925a2\"], \"transaction_hash\": \"0x227f3d9d5ce6680bdf2991576c1a90aca8184ca26055bae92d16c58e3e13340\", \"type\": \"DEPLOY\"}, {\"contract_address\": \"0x6ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae\", \"entry_point_selector\": \"0x3d7905601c217734671143d457f0db37f7f8883112abd34b92c4abfeafde0c3\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x1e2cd4b3588e8f6f9c4e89fb0e293bf92018c96d7a93ee367d29a284223b6ff\", \"0x71d1e9d188c784a0bde95c1d508877a0d93e9102b37213d1e13f3ebc54a7751\"], \"signature\": [], \"transaction_hash\": \"0x376ff82431b52ca1fbc4942de80bc1b01d8e5cd1eeab5a277b601b510f2cab2\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x31c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52\", \"contract_address_salt\": \"0x63d1a6f8130958509e2e695c25b147f43f66f56bba49fddb7ee363d8f57a774\", \"class_hash\": \"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8\", \"constructor_calldata\": [\"0xdf28e613c065616a2e79ca72f9c1908e17b8c913972a9993da77588dc9cae9\", \"0x1432126ac23c7028200e443169c2286f99cdb5a7bf22e607bcd724efa059040\"], \"transaction_hash\": \"0x25f20c74821d84f62989a71fceef08c967837b63bae31b279a11343f10d874a\", \"type\": \"DEPLOY\"}, {\"contract_address\": \"0x31c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52\", \"entry_point_selector\": \"0x218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x735596016a37ee972c42adef6a3cf628c19bb3794369c65d2c82ba034aecf2c\", \"0x1\"], \"signature\": [], \"transaction_hash\": \"0x2d10272a8ba726793fd15aa23a1e3c42447d7483ebb0b49df8b987590fe0055\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x735596016a37ee972c42adef6a3cf628c19bb3794369c65d2c82ba034aecf2c\", \"entry_point_selector\": \"0x218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x31c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52\", \"0x0\"], \"signature\": [], \"transaction_hash\": \"0xb05ba5cd0b9e0464d2c1790ad93a159c6ef0594513758bca9111e74e4099d4\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x6ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae\", \"entry_point_selector\": \"0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x1a7cf8b8027ec2d8fd04f1277f3f8ae6379ca957c5fec9ee7f59d56d86a26e4\", \"0x2\", \"0x28dff6722aa73281b2cf84cac09950b71fa90512db294d2042119abdd9f4b87\", \"0x57a8f8a019ccab5bfc6ff86c96b1392257abb8d5d110c01d326b94247af161c\"], \"signature\": [], \"transaction_hash\": \"0x4d16393d940fb4a97f20b9034e2a5e954201fee827b2b5c6daa38ec272e7c9c\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x31c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52\", \"entry_point_selector\": \"0x19a35a6e95cb7a3318dbb244f20975a1cd8587cc6b5259f15f61d7beb7ee43b\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x6ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae\", \"0x5f750dc13ed239fa6fc43ff6e10ae9125a33bd05ec034fc3bb4dd168df3505f\"], \"signature\": [], \"transaction_hash\": \"0x9e80672edd4927a79f5384e656416b066f8ef58238227ac0fcea01952b70b5\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x6ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae\", \"entry_point_selector\": \"0x3d7905601c217734671143d457f0db37f7f8883112abd34b92c4abfeafde0c3\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x449908c349e90f81ab13042b1e49dc251eb6e3e51092d9a40f86859f7f415b0\", \"0x2670b3a8266d5046696a4b79f7433d117d3a19166f15bbd8585822c4e9b7491\"], \"signature\": [], \"transaction_hash\": \"0x387b5b63e40d4426754895fe52adf668cf8fde2a02aa9b6d761873f31af3462\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}, {\"contract_address\": \"0x6ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae\", \"entry_point_selector\": \"0x3d7905601c217734671143d457f0db37f7f8883112abd34b92c4abfeafde0c3\", \"entry_point_type\": \"EXTERNAL\", \"calldata\": [\"0x449908c349e90f81ab13042b1e49dc251eb6e3e51092d9a40f86859f7f415b0\", \"0x6cb6104279e754967a721b52bcf5be525fdc11fa6db6ef5c3a4db832acf7804\"], \"signature\": [], \"transaction_hash\": \"0x4f0cdff0d72fc758413a16db2bc7580dfec7889a8b921f0fe08641fa265e997\", \"max_fee\": \"0x0\", \"type\": \"INVOKE_FUNCTION\"}], \"timestamp\": 1637069048, \"sequencer_address\": \"0x37b2cd6baaa515f520383bee7b7094f892f4c770695fc329a8973e841a971ae\", \"transaction_receipts\": [{\"transaction_index\": 0, \"transaction_hash\": \"0xe0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 29, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 1, \"transaction_hash\": \"0x12c96ae3c050771689eb261c9bf78fac2580708c7f1f3d69a9647d8be59f1e1\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 29, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 2, \"transaction_hash\": \"0xce54bbc5647e1c1ea4276c01a708523f740db0ff5474c77734f73beec2624\", \"l2_to_l1_messages\": [{\"from_address\": \"0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6\", \"to_address\": \"0xC84DD7fd43a7deFb5b7a15c4FBBE11cBba6dB1bA\", \"payload\": [\"0xc\", \"0x22\"]}], \"events\": [], \"execution_resources\": {\"n_steps\": 31, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 3, \"transaction_hash\": \"0x1c924916a84ef42a3d25d29c5d1085fe212de04feadc6e88d4c7a6e5b9039bf\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 238, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 4, \"transaction_hash\": \"0xa66c346e273cc49510ef2e1620a1a7922135cb86ab227b86e0afd12243bd90\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 165, \"builtin_instance_counter\": {\"pedersen_builtin\": 2, \"range_check_builtin\": 7, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 22}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 5, \"transaction_hash\": \"0x5c71675616b49fb9d16cac8beaaa65f62dc5a532e92785055c15c825166dbbf\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 209, \"builtin_instance_counter\": {\"pedersen_builtin\": 2, \"range_check_builtin\": 8, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 24}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 6, \"transaction_hash\": \"0x60e05c41a6622592a2e2eff90a9f2e495296a3be9596e7bc4dfbafce00d7a6a\", \"l2_to_l1_messages\": [{\"from_address\": \"0x31c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280\", \"to_address\": \"0x0000000000000000000000000000000000000001\", \"payload\": [\"0xc\", \"0x22\"]}], \"events\": [], \"execution_resources\": {\"n_steps\": 332, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 7, \"transaction_hash\": \"0x5634f2847140263ba59480ad4781dacc9991d0365145489b27a198ebed2f969\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 178, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 8, \"transaction_hash\": \"0xb049c384cf75174150a2540835cc2abdcca1d3a3750298a1741a621983e35a\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 29, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 9, \"transaction_hash\": \"0x227f3d9d5ce6680bdf2991576c1a90aca8184ca26055bae92d16c58e3e13340\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 29, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 10, \"transaction_hash\": \"0x376ff82431b52ca1fbc4942de80bc1b01d8e5cd1eeab5a277b601b510f2cab2\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 25, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 11, \"transaction_hash\": \"0x25f20c74821d84f62989a71fceef08c967837b63bae31b279a11343f10d874a\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 29, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 12, \"transaction_hash\": \"0x2d10272a8ba726793fd15aa23a1e3c42447d7483ebb0b49df8b987590fe0055\", \"l2_to_l1_messages\": [{\"from_address\": \"0x31c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52\", \"to_address\": \"0x0000000000000000000000000000000000000001\", \"payload\": [\"0xc\", \"0x22\"]}], \"events\": [], \"execution_resources\": {\"n_steps\": 332, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 13, \"transaction_hash\": \"0xb05ba5cd0b9e0464d2c1790ad93a159c6ef0594513758bca9111e74e4099d4\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 238, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 14, \"transaction_hash\": \"0x4d16393d940fb4a97f20b9034e2a5e954201fee827b2b5c6daa38ec272e7c9c\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 169, \"builtin_instance_counter\": {\"pedersen_builtin\": 2, \"range_check_builtin\": 7, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 20}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 15, \"transaction_hash\": \"0x9e80672edd4927a79f5384e656416b066f8ef58238227ac0fcea01952b70b5\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 178, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 16, \"transaction_hash\": \"0x387b5b63e40d4426754895fe52adf668cf8fde2a02aa9b6d761873f31af3462\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 25, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}, {\"transaction_index\": 17, \"transaction_hash\": \"0x4f0cdff0d72fc758413a16db2bc7580dfec7889a8b921f0fe08641fa265e997\", \"l2_to_l1_messages\": [], \"events\": [], \"execution_resources\": {\"n_steps\": 25, \"builtin_instance_counter\": {\"pedersen_builtin\": 0, \"range_check_builtin\": 0, \"bitwise_builtin\": 0, \"output_builtin\": 0, \"ecdsa_builtin\": 0, \"ec_op_builtin\": 0}, \"n_memory_holes\": 0}, \"actual_fee\": \"0x0\"}]}"
	bjson := []byte(rjson)
	json.Unmarshal(bjson, &block)
	return &block
}
