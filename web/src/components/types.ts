

export interface Param {
    route: string[]; // 路由参数
    req_datas: ReqData[]; // 请求参数，键值对形式
    res_code: number; // 响应码
    res_content_type: string; // 响应类型
    res_data: string; // 响应内容
}

interface ReqData {
    key: string;
    value: string;
}

export interface Form {
    id: string;
    name: string; // 名称
    group: string; // 分组
    methods: string[]; // 方法
    req_content_type: string; // 请求类型
    path: string; // 路径
    description: string; // 描述
    params: Param[]; // 参数列表
}

