export interface AsApplicationTableColumns {    
    id:number;  // ID    
    fileUrl:string;  // 文件    
    name:string;  // 应用名称    
    bundleId:string;  // 包名    
    version:string;  // 版本号    
    size:number;  // 文件大小
    showSize:string;  // 展示文件大小
    type:number;  // 类型
    description:string;  // 描述    
    iconUrl:string;  // 图标    
    lock:number;  // 付费    
    lanzou:number;  // 是否蓝奏云链接    
    weigh:number;  // 权重    
    active:number;  // 是否启用    
    note:string;  // 备注    
    createdAt:string;  // 创建时间    
}


export interface AsApplicationInfoData {    
    id:number|undefined;        // ID    
    fileUrl:string|undefined; // 文件    
    name:string|undefined; // 应用名称    
    bundleId:string|undefined; // 包名    
    version:string|undefined; // 版本号    
    size:number|undefined; // 文件大小
    showSize:string|undefined; // 文件大小
    type:number|undefined; // 类型
    description:string|undefined; // 描述    
    iconUrl:string|undefined; // 图标    
    lock:boolean; // 付费    
    lanzou:number|undefined; // 是否蓝奏云链接    
    weigh:number|undefined; // 权重    
    active:boolean; // 是否启用    
    note:string|undefined; // 备注    
    createdBy:number|undefined; // 创建人    
    updatedBy:number|undefined; // 修改人    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 修改时间    
}


export interface AsApplicationTableDataState {
    ids:any[];
    tableData: {
        data: Array<AsApplicationTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            fileUrl: string|undefined;            
            name: string|undefined;            
            bundleId: string|undefined;            
            version: string|undefined;            
            size: number|undefined;
            showSize: string|undefined;
            type: number|undefined;
            description: string|undefined;            
            iconUrl: string|undefined;            
            lock: number|undefined;            
            lanzou: number|undefined;            
            weigh: number|undefined;            
            active: number|undefined;            
            note: string|undefined;            
            createdBy: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
    sourceImportMessage: string;
}


export interface AsApplicationEditState{
    loading:boolean;
    isShowDialog: boolean;
    appDownloadMessage: string;
    formData:AsApplicationInfoData;
    rules: object;
}