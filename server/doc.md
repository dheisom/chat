# Documentation

This is a tiny documentation about this API

## Create user

To create a user with the version 1 of the API you have to make a POST request
to `/api/v1/createUser` with the some informations:

|  JSON      |  Type  |  Description                                 |  Default  |
| :--------- | :----- | :------------------------------------------- | :-------- |
| `user`     | Text   | The name that will be showed on chat title   | *         |
| `username` | Text   | Username to found you with the search method | Random    |
| `bio`      | Text   | A tiny description about you                 | Empty     |

_Info_: All fields marked with `*` are mandatory!

## Methods

To call a method send a request to: `/api/v{versionNumber}/{yourToken}/{methodName}`

## Version: 1

### **GET**: `getMessages`

Get received messages

|  Query  |  Type    |  Description                                |  Default  |
| :------ | :------- | :------------------------------------------ | :-------- |
| `seek`  | Integer  | The position since the messages will be get | 1         |
| `limit` | Integer  | Limit of messages to be in the response     | 50        |

### **POST**: `sendMessage`

Send a message to a user

|  JSON     |  Type                    |  Description                                |
| :-------- | :----------------------- | :------------------------------------------ |
| `text`    | Text                     | The message text to be send                 |
| `to_user` | Unsigned integer or Text | The user ID or username of the other client |

_Info_: All fields are mandatory

### **GET**: `getMe`

Get informations about me

### **GET**: `getUser`

Get informations about another user

|  Query  |  Type                    |  Description                             |
| :------ | :----------------------- | :--------------------------------------- |
| `user`  | Unsigned integer or Text | User ID or username of the other client  |

## **GET**: `getChats`

Get chats that you have sended or received messages
