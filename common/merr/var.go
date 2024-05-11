package merr

// OK 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码000000-099999

const GLOBAL_ERROR uint32 = 000000
const SERVER_COMMON_ERROR uint32 = 000001
const REUQEST_PARAM_ERROR uint32 = 000002
const TOKEN_EXPIRE_ERROR uint32 = 000003
const TOKEN_GENERATE_ERROR uint32 = 000004
const DB_ERROR uint32 = 000005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 000006

// 用户10000-19999

// 课程20000-29999
