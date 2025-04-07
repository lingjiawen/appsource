export interface AsRedeemCodeTableColumns {
    id:number;  // ID
    code:string;  // 兑换码    
    udid:string;  // 设备码    
    type:number;  // 类型    
    active:number;  // 是否激活    
    activeAt:string;  // 激活时间    
    createdAt:string;  // 生成时间    
}


export interface AsRedeemCodeInfoData {    
    id:number|undefined;        // ID    
    code:string|undefined; // 兑换码    
    udid:string|undefined; // 设备码    
    type:number|undefined; // 类型
    active:boolean; // 是否激活    
    activeAt:string|undefined; // 激活时间    
    createdBy:number|undefined; // 创建人    
    updatedBy:number|undefined; // 修改人    
    createdAt:string|undefined; // 生成时间    
    updatedAt:string|undefined; // 修改时间    
    deletedAt:string|undefined; // 删除时间    
}

export interface AsRedeemCodeAddData {
    id:number|undefined;        // ID
    prefix:string|undefined; // 前缀
    quantity:number|undefined; // 数量
    type:number|undefined; // 类型
}

export interface AsRedeemCodeTableDataState {
    ids:any[];
    tableData: {
        data: Array<AsRedeemCodeTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            code: string|undefined;            
            udid: string|undefined;            
            type: number|undefined;            
            active: number|undefined;            
            activeAt: string|undefined;            
            createdBy: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface AsRedeemCodeEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:AsRedeemCodeInfoData;
    rules: object;
}

export interface AsRedeemCodeAddState{
    loading:boolean;
    isShowDialog: boolean;
    formData:AsRedeemCodeAddData;
    rules: object;
}