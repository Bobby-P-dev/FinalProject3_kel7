# Dokumentasi Final Project 3

Link github = https://github.com/Bobby-P-dev/FinalProject3_kel7

# dokumentasi API

Domain Api = https://final3kel7.adaptable.app


Domain Database = postgres://otuexgqo:JLlM-Y3kDSk3d0tY_sO58kWZKYynZAlC@flora.db.elephantsql.com/otuexgqo

## account admin
    {
    "email":    "admin@gmail.com",
    "password": "123456"
    }

***

# user
## regist account

Method = POST

Domain = https://final3kel7.adaptable.app/user/register


request body 
```
{
    "full_name": string,
    "email":    string,
    "password": string
}
```

## Login

Method = POST

Domain = https://final3kel7.adaptable.app/user/login

request body

```
{
    "email":    string,
    "password": string
}
```

## put account

Method = PUT

Domain = https://final3kel7.adaptable.app/user/update-account

request body
```
{
    "full_name": string,
    "email":    string
}
```

## delete account

Method = DELETE

Domain = https://final3kel7.adaptable.app/user/delete-account

request
`
bearer token authorizaiton
`
***

# category
## create category

Method = POST

Domain = https://final3kel7.adaptable.app/category/post

request
`
bearer token authorizaiton only admin
`

request body
```
  {
      "type":   string
  }
```

## get category

Method = GET

Domain = https://final3kel7.adaptable.app/category/get

request
`
bearer token authorization
`

## patch category

Method = PATCH

Domain = https://final3kel7.adaptable.app/category/patch/id

request
`
id param(int) & bearer token authorization only admin
`

request body
```
  {
    "type": string
  }
```

## delete category

Method = DELETE

Domain = https://final3kel7.adaptable.app/category/delete/id

request
`
id param(int) & bearer token authorization only admin
`


----------------------------------------------------------------------------

# task
## create task

Method = POST

Domain = https://final3kel7.adaptable.app/tasks/post

request
`
bearer token authorization
`

request body
```
{
    "title":  sting
    "description":  string,
    "category_id":  int
}
```

## get task

Method = GET

Domain = https://final3kel7.adaptable.app/tasks/get

request
`
bearer token authorization
`

## put task

Method = PUT

Domain = https://final3kel7.adaptable.app/tasks/put/id

request
`
id param(int) & bearer token authorization
`

request body
```
    {
      "title":   string,
      "description":   string
    }
```

## patch task status

Method = PATCH

Domain = https://final3kel7.adaptable.app/tasks/patch-status/id

request
`
id param(int) & bearer token authorization
`

request body
```
    {
      "status": boolean = true / false
    }
```

## patch task categoryid

Method = PATCH

Domain = https://final3kel7.adaptable.app/tasks/patch-category/id

request
`
id param(int) & bearer token authorization
`

request body
```
  {
     "category_id": int
  }
```

## delete task

Method = DELETE

Domain = https://final3kel7.adaptable.app/tasks/delete/id

request
`
id param(int) & bearer token authorization
`