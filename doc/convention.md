# 约定

## API返回格式

### 所有列表返回格式:  
{
    "status_code": 200,
    "msg": "查询成功",
    data: []
}

### 所有单个查询返回格式:  
{
    "status_code": 200,
    "msg": "查询成功",
    data: {
      .... 字段列表, 找不到返回空对象
    }
}

### 参数缺失，参数不合法(包括不限于转换到int失败, email 或电话格式不对):  
{
    "status_code": 400,
    "msg": "查询成功",
}

## 数据库字段

表名都用复数, 模型名用单数  

### users 表

gender 1男 0女  
role 专家1，卖家2，其他普通用户

### chips 表

manufacture_date 制造日期 精确到年月就可以了

### buy_requests 表

amount -1为不限数量  

### 其他表  
is_verified 暂时为方便调试，默认通过审核  
暂时没有写Authentication Middleware，需要传user_id 以识别用户身份  
