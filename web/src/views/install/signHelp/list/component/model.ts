export interface SignHelpTableColumns {    
    id:number;  // ID    
    title:string;  // 标题    
    content:string;  // 内容    
    isExpand:number;  // 默认展开    
    weigh:number;  // 权重    
    createdAt:string;  // 创建日期    
}


export interface SignHelpInfoData {    
    id:number|undefined;        // ID    
    title:string|undefined; // 标题    
    content:string|undefined; // 内容    
    isExpand:boolean; // 默认展开    
    weigh:number|undefined; // 权重    
    createdAt:string|undefined; // 创建日期    
    updatedAt:string|undefined; // 修改日期    
}


export interface SignHelpTableDataState {
    ids:any[];
    tableData: {
        data: Array<SignHelpTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            title: string|undefined;            
            content: string|undefined;            
            isExpand: number|undefined;            
            weigh: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface SignHelpEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SignHelpInfoData;
    rules: object;
}