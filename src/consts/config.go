package consts

var DefaultConfig =
`[daemon]
port = "8080"

[default]
visible = true
user_list = false
user_max = 2
filter_token = [
    "whatever",
    "stupid",
]
filter_whole = [
    "hate you",
    "kill"
]

[session.abc]
user_list = true
connection_max = 30
`
