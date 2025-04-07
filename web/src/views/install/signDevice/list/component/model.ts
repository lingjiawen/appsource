export interface SignDeviceTableColumns {    
    id:number;  // ID    
    udid:string;  // 设备码    
    name:string;  // 证书名    
    certId:string;  // 证书标识    
    addTime:number;  // 添加时间    
    model:string;  // 设备型号    
    expireTime:number;  // 过期时间    
    redeemCode:string;  // 兑换卡密    
    accountType:number;  // 时效类型    
    warrantyType:number;  // 售后类型    
    deviceType:string;  // 设备类型    
    status:string;  // 状态    
    pool:number;  // 出书方式    
    apiPlatform:number;  // 对接平台    
    apiWarrantyType:number;  // 对接售后类型    
    active:number;  // 禁用    
    createdBy:string;  // 创建人    
    createdAt:string;  // 创建时间    
}


export interface SignDeviceInfoData {    
    id:number|undefined;        // ID    
    udid:string|undefined; // 设备码    
    name:string|undefined; // 证书名    
    certId:string|undefined; // 证书标识    
    deviceId:string|undefined; // 设备标识    
    p12:string|undefined; // 证书文件    
    mobileprovision:string|undefined; // 描述文件    
    addTime:number|undefined; // 添加时间    
    p12Password:string|undefined; // 证书密码    
    model:string|undefined; // 设备型号    
    expireTime:number|undefined; // 过期时间    
    serialNumber:string|undefined; // 序列号    
    redeemCode:string|undefined; // 兑换卡密    
    accountType:number|undefined; // 时效类型    
    warrantyType:number|undefined; // 售后类型    
    deviceType:string|undefined; // 设备类型    
    status:string|undefined; // 状态    
    pool:number|undefined; // 出书方式    
    apiPlatform:number|undefined; // 对接平台    
    apiWarrantyType:number|undefined; // 对接售后类型    
    active:boolean; // 禁用    
    createdBy:number|undefined; // 创建人    
    updatedBy:number|undefined; // 修改人    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 修改时间    
    deletedAt:string|undefined; // 删除时间    
}


export interface SignDeviceTableDataState {
    ids:any[];
    tableData: {
        data: Array<SignDeviceTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            udid: string|undefined;            
            name: string|undefined;            
            certId: string|undefined;            
            deviceId: string|undefined;            
            addTime: number|undefined;            
            model: string|undefined;            
            expireTime: number|undefined;            
            redeemCode: string|undefined;            
            accountType: number|undefined;            
            warrantyType: number|undefined;            
            deviceType: string|undefined;            
            status: string|undefined;            
            pool: number|undefined;            
            apiPlatform: number|undefined;            
            apiWarrantyType: number|undefined;            
            active: number|undefined;            
            createdBy: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface SignDeviceEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SignDeviceInfoData;
    rules: object;
}