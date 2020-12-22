package api

type Signinwithpasswd struct {
	Uid     string `yaml:"uid,omitempty" json:"uid,omitempty"`
	User_id int64  `yaml:"user_id,omitempty" json:"user_id,omitempty"`
}

//pic: jiami(id,secret_str),secret_str
