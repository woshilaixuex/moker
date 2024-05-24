package merr

// OK 成功返回
const OK uint32 = 200

// UNKNOWN_ERROR 未知错误/未定义
const UNKNOWN_ERROR uint32 = 999999

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码000000-000999

const GLOBAL_ERROR uint32 = 000000

const SERVER_TIME_OUT uint32 = 000001

// 用户中心001000-001999

const USER_ERROR = 001000
const VERCODE_ERROR = 001001
