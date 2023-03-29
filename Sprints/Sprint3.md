--Detail work you've completed in Sprint 3

Connected go backend to database that can use and handle SQL queries

Established working login

Created a password authentication function for the login page that encrypts and decrypts user passwords making them undetectable to us as admins and any potential threats on security

--List frontend unit tests


--List backend unit tests
- Tests Register function located in authController.go that registers a user to the database
- Tests CreateUser function located in userController.go which ensures that a user is successfully created

--Show updated documentation for your backend API 

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
- Paginated users to limit the number of users per page
    - Keeps track of total pages, current, and last page
- Created functions to allow user email/password to be updated
- Product (food items) controller that allows addition of items to database
    - Includes title, description, image, and price (tentative)

