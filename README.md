| HTTP Metodu | Endpoint                   | İşlev                                                                               |
|-------------|----------------------------|-------------------------------------------------------------------------------------|
| GET         | /users                     | restaurantmanagement/controllers.GetUsers.func1 (3 handlers)                        |
| GET         | /users/:user_id            | restaurantmanagement/controllers.GetUser.func1 (3 handlers)                         |
| POST        | /users/signup              | restaurantmanagement/controllers.SignUp.func1 (3 handlers)                       |
| POST        | /users/login               | restaurantmanagement/controllers.Login.func1 (3 handlers)                       |
|             |                            |                                                                                      |
| GET         | /foods                     | restaurantmanagement/controllers.GetFoods.func1 (3 handlers)                        |
| GET         | /foods/:food_id            | restaurantmanagement/controllers.GetFood.func1 (3 handlers)                         |
| POST        | /foods/                    | restaurantmanagement/controllers.CreateFood.func1 (3 handlers)                       |
| PATCH       | /foods/:food_id            | restaurantmanagement/controllers.UpdateFood.func1 (3 handlers)                       |
|             |                            |                                                                                      |
| GET         | /menus                     | restaurantmanagement/controllers.GetMenus.func1 (3 handlers)                        |
| GET         | /menus/:menu_id            | restaurantmanagement/controllers.GetMenu.func1 (3 handlers)                         |
| POST        | /menus/                    | restaurantmanagement/controllers.CreateMenu.func1 (3 handlers)                       |
| PATCH       | /menus/:menu_id            | restaurantmanagement/controllers.UpdateMenu.func1 (3 handlers)                       |
|             |                            |                                                                                      |
| GET         | /tables                    | restaurantmanagement/controllers.GetTables.func1 (3 handlers)                       |
| GET         | /tables/:table_id          | restaurantmanagement/controllers.GetTable.func1 (3 handlers)                        |
| POST        | /tables/                   | restaurantmanagement/controllers.CreateTable.func1 (3 handlers)                      |
| PATCH       | /tables/:table_id          | restaurantmanagement/controllers.UpdateTable.func1 (3 handlers)                      |
|             |                            |                                                                                      |
| GET         | /orders                    | restaurantmanagement/controllers.GetOrders.func1 (3 handlers)                        |
| GET         | /orders/:order_id          | restaurantmanagement/controllers.GetOrder.func1 (3 handlers)                         |
| POST        | /orders/                   | restaurantmanagement/controllers.CreateOrder.func1 (3 handlers)                       |
| PATCH       | /orders/:order_id          | restaurantmanagement/controllers.UpdateOrder.func1 (3 handlers)                       |
|             |                            |                                                                                      |
| GET         | /orderitems                | restaurantmanagement/controllers.GetOrderItems.func1 (3 handlers)                    |
| GET         | /orderitems/:orderitem_id  | restaurantmanagement/controllers.GetOrderItem.func1 (3 handlers)                      |
| GET         | /orderitems-order/:order_id | restaurantmanagement/controllers.GetOrderItemsByOrder.func1 (3 handlers)              |
| POST        | /orderitems/               | restaurantmanagement/controllers.CreateOrderItem.func1 (3 handlers)                    |
| PATCH       | /orderitems/:orderitem_id   | restaurantmanagement/controllers.UpdateOrderItem.func1 (3 handlers)                    |
