export interface SignPlatformTableColumns {    
    id:number;  // ID    
    name:string;  // 平台名    
    code:string;  // 平台标识    
    baseUrl:string;  // 域名    
    openSsl:number;  // 开启SSL    
    status:number;  // 启用    
    token:string;  // 对接Token    
    createdAt:string;  // 创建时间    
    weigh:number;  // 权重    
}


export interface SignPlatformInfoData {    
    id:number|undefined;        // ID    
    name:string|undefined; // 平台名    
    code:string|undefined; // 平台标识    
    baseUrl:string|undefined; // 域名    
    openSsl:boolean; // 开启SSL    
    status:boolean; // 启用    
    token:string|undefined; // 对接Token    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 修改时间    
    weigh:number|undefined; // 权重    
}


export interface SignPlatformTableDataState {
    ids:any[];
    tableData: {
        data: Array<SignPlatformTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            name: string|undefined;            
            code: string|undefined;            
            baseUrl: string|undefined;            
            openSsl: number|undefined;            
            status: number|undefined;            
            token: string|undefined;            
            createdAt: string|undefined;            
            weigh: number|undefined;            
            dateRange: string[];
        };
    };
}


export interface SignPlatformEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SignPlatformInfoData;
    rules: object;
}