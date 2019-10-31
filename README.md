# URLSPACE
A place to store links to your social media accounts or any accounts at all.
No authentication mechanism set in place yet.   API is live at https://urlspace.herokuapp.com

## SETUP
Environment is linux
```
git clone https://github.com/iambenkay/urlspace.git
cd urlspace
yarn
npm start
```

## HOW TO USE
### Create a new URL space with your username
```js
POST /api/v1/:user

body:
{
    "name": String
    "link": String
}
```
The request above creates searches for an existing URL space belonging to `user` and creates one if there are none. It then adds the body to the URL space and returns the content of the URL space.  
  
### Fetch contents of a URL space
```js
GET /api/v1/:user
```
The request above fetches the content of the URL space belonging to `user`
  
### Fetch URL of a particular account in URL space
```js
GET /api/v1/:user/:account
```
The request above fetches the url of the particular `account` in the URL space belonging to `user`
  
Now you can visit `/:user/:account` and it will redirect you to the link of the account.
