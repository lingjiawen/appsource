export interface SysConfigGroupTableColumns {    
    id:number;  // 主键    
    groupName:string;  // 配置组名称    
    groupKey:string;  // 配置组键名    
    createdAt:string;  // 创建时间    
}


export interface SysConfigGroupInfoData {    
    id:number|undefined;        // 主键    
    groupName:string|undefined; // 配置组名称    
    groupKey:string|undefined; // 配置组键名    
    createBy:number|undefined; // 创建者    
    updateBy:number|undefined; // 更新者    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 修改时间    
}


export interface SysConfigGroupTableDataState {
    ids:any[];
    tableData: {
        data: Array<SysConfigGroupTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            groupName: string|undefined;            
            groupKey: string|undefined;            
            createBy: number|undefined;            
            updateBy: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface SysConfigGroupEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SysConfigGroupInfoData;
    rules: object;
}