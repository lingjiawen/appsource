export interface SignRedeemCodeTableColumns {    
    id:number;  // ID    
    code:string;  // 兑换码    
    udid:string;  // 设备码    
    certId:string;  // 证书标识    
    accountType:number;  // 时效类型    
    warrantyType:number;  // 售后类型    
    deviceType:string;  // 设备类型    
    pool:number;  // 出书方式    
    apiPlatform:number;  // 对接平台
    force:number;  // 强制
    note:string;  // 备注    
    apiWarrantyType:number;  // 对接售后类型    
    banned:number;  // 禁用    
    active:number;  // 激活    
    activeAt:string;  // 激活时间    
    createdAt:string;  // 创建时间    
}


export interface SignRedeemCodeInfoData {    
    id:number|undefined;        // ID    
    code:string|undefined; // 兑换码    
    udid:string|undefined; // 设备码    
    certId:string|undefined; // 证书标识    
    accountType:number|undefined; // 时效类型    
    warrantyType:number|undefined; // 售后类型    
    deviceType:string|undefined; // 设备类型    
    pool:number|undefined; // 出书方式    
    apiPlatform:number|undefined; // 对接平台    
    note:string|undefined; // 备注    
    apiWarrantyType:number|undefined; // 对接售后类型
    force:boolean; // 强制
    banned:boolean; // 禁用    
    active:boolean; // 激活    
    activeAt:string|undefined; // 激活时间    
    createdBy:number|undefined; // 创建人    
    updatedBy:number|undefined; // 修改人    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 修改时间    
    deletedAt:string|undefined; // 删除时间    
}

export interface SignRedeemCodeAddData {
    id:number|undefined;        // ID
    prefix:string|undefined; // 前缀
    quantity:number|undefined; // 数量
    accountType:number|undefined; // 账号类型
    warrantyType:number|undefined; // 质保类型
    deviceType:string|undefined; // 设备类型
    pool:number|undefined; // 出书方式
    force:number|undefined; // 强制
    apiPlatform:number|undefined; // 对接平台
    apiWarrantyType:number|undefined; // 对接质保
    note:string|undefined; // 备注
}

export interface SignRedeemCodeTableDataState {
    ids:any[];
    tableData: {
        data: Array<SignRedeemCodeTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            code: string|undefined;            
            udid: string|undefined;            
            certId: string|undefined;            
            accountType: number|undefined;            
            warrantyType: number|undefined;            
            deviceType: string|undefined;            
            pool: number|undefined;            
            apiPlatform: number|undefined;            
            apiWarrantyType: number|undefined;            
            note: string|undefined;
            force: number|undefined;
            banned: number|undefined;            
            active: number|undefined;            
            activeAt: string|undefined;            
            createdBy: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface SignRedeemCodeEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SignRedeemCodeInfoData;
    rules: object;
}

export interface SignRedeemCodeAddState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SignRedeemCodeAddData;
    rules: object;
}