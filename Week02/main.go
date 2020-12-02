package Week02


func main(){

	user := UserService{}
	name := "test"

	// 如果这个错误是我在意的，根据错误只处理一次的原则，底层已经wrap，那么作为上层，我直接抛上去。
	_, err := user.GetNameWithMessage(name)
	if err != nil {
		// 返回给上层 return error
	}

	// 如果这个错误业务上来讲是我不在意的，那么到我这一层去处理逻辑即可。
	u,_ := user.GetNameNotCare(name)
	if u != nil {
		// 业务处理
	}
}

