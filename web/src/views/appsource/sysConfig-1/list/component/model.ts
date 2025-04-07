export interface SysConfigTableColumns {    
    configId:number;  // 参数主键    
    configName:string;  // 参数名称    
    configKey:string;  // 参数键名    
    configValue:string;  // 参数键值    
    configType:number;  // 系统内置（Y是 N否）    
    group:string;  // 参数分组    
    linkedGroup?:LinkedSysConfigSysConfigGroup; // 参数分组    
    remark:string;  // 备注    
    createdAt:string;  // 创建时间    
    linkedSysConfigSysConfigGroup:LinkedSysConfigSysConfigGroup;    
}


export interface SysConfigInfoData {    
    configId:number|undefined;        // 参数主键    
    configName:string|undefined; // 参数名称    
    configKey:string|undefined; // 参数键名    
    configValue:string|undefined; // 参数键值    
    configType:number|undefined; // 系统内置（Y是 N否）    
    group:string|undefined; // 参数分组    
    linkedGroup?:LinkedSysConfigSysConfigGroup; // 参数分组    
    createBy:number|undefined; // 创建者    
    updateBy:number|undefined; // 更新者    
    remark:string|undefined; // 备注    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 修改时间    
    linkedSysConfigSysConfigGroup?:LinkedSysConfigSysConfigGroup;    
}


export interface LinkedSysConfigSysConfigGroup {	
    groupKey:string|undefined;    // 配置组键名	
    groupName:string|undefined;    // 配置组名称	
}


export interface SysConfigTableDataState {
    configIds:any[];
    tableData: {
        data: Array<SysConfigTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            configId: number|undefined;            
            configName: string|undefined;            
            configKey: string|undefined;            
            configValue: string|undefined;            
            configType: number|undefined;            
            group: string|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface SysConfigEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SysConfigInfoData;
    rules: object;
}