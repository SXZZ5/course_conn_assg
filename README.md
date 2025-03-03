## Usage
- Click on **Old Cookies** to view all cookies. If you are not logged in, No cookies will be shown. 
- **Signup:** From the home page, click on signup and set a username, password. If successful you should be redirected to page showing you cookies.
- **Logout:** Click on logout button on the home page. 
- To create a new cookie before expiry, Go back to the home page and click on login and provide the username and password for your account. This should trigger creation of a new cookie for you. 
- Successful login should redirect to the "Old Cookies" page automatically.

## Running
- In client directory
  - npm install
  - npm run dev
- In backend directory
  - there should be a MySQL service running on the OS on port 3306 with a database named "regsys". TODO: replace these with environment variables. 
  - go run schema/setup.go
  - go run main.go
- Visit localhost:5173/ in the browser




