export interface SignAboutTableColumns {    
    id:number;  // ID    
    icon:string;  // 图标    
    title:string;  // 标题    
    subtitle:string;  // 内容    
    isLink:number;  // 是否链接    
    url:string;  // 链接    
    group:number;  // 分组    
    weigh:number;  // 权重    
    createdAt:string;  // 创建日期    
}


export interface SignAboutInfoData {    
    id:number|undefined;        // ID    
    icon:string|undefined; // 图标    
    title:string|undefined; // 标题    
    subtitle:string|undefined; // 内容    
    isLink:boolean; // 是否链接    
    url:string|undefined; // 链接    
    group:number|undefined; // 分组    
    weigh:number|undefined; // 权重    
    createdAt:string|undefined; // 创建日期    
    updatedAt:string|undefined; // 修改日期    
}


export interface SignAboutTableDataState {
    ids:any[];
    tableData: {
        data: Array<SignAboutTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            icon: string|undefined;            
            title: string|undefined;            
            subtitle: string|undefined;            
            isLink: number|undefined;            
            url: string|undefined;            
            group: number|undefined;            
            weigh: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface SignAboutEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SignAboutInfoData;
    rules: object;
}