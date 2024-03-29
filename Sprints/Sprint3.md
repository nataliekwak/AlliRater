**Detail work you've completed in Sprint 3**


Connected go backend to database that can use and handle SQL queries

Established working login

Created a password authentication function for the login page that encrypts and decrypts user passwords making them undetectable to us as admins and any potential threats on security



**List frontend unit tests**


-The designated homepage will be examined for the presence and visibility of the dashboard button. An assessment of the website's layout will be conducted to ensure that the button is prominently displayed and easily identifiable

-Upon confirming the presence of the dashboard button, it will be tested for functionality by clicking on it. The dashboard page will be thoroughly analyzed to ensure that any updates made to the homepage are promptly displayed, as expected. It will verify that the updated dashboard page’s url is ‘/#’

-The login page will be accessed to confirm its presence and accessibility. It will be verified that the login page's URL is '/login'

-The register page will be accessed to confirm its presence and accessibility. It will be verified that the register page's URL is '/register'


**List backend unit tests**


- Tests Register function located in authController.go that registers a user to the database
- Tests CreateUser function located in userController.go which ensures that a user is successfully created



**Show updated documentation for your backend API**


Sprint 3 Backend Documentation:
- Implemented fiber framework to run server
    - Connected to database using MySQL Workbench and GORM
- Created and organized packages into folders
    - Includes controllers, databases, models, and routes
- Implemented Postman to send requests
- Created users, permission, and role models
    - Register users function
    - Users struct to store each user's information
    - Roles model creates new table in database with role information
    - Permissions model gives functions to view and edit users/items
- Created entity, order, paginate, permission models
    - Gives information about food item (quantity, name, etc.)
- Installed go realize package to add listener every time changes are made
- Added migrations using GORM
- Hashed password with bcrypt go package
- Login Function
    - Uses post requests
    - Sends message if user not found or incorrect password
- Created/generated JWT token with secret key
    - Stored token into cookie
- Added logout function
    - Created cookie and post request
- Implemented middleware functions to increase efficiency
    - Handles permissions allowing for view and edit access
    - Secures routes with authorization for certain users
- Paginated users to limit the number of users per page
    - Keeps track of total pages, current, and last page
- Created functions to allow user email/password to be updated
- Product (food items) controller that allows addition of items to database
    - Includes title, description, image, and price (tentative)
- Created function that exports data as a CSV file
    - Includes name, email, item, price, quantity
- Executed SQL query to format chart with database information by date