--Detail work you've completed in Sprint 3

Connected go backend to database that can use and handle SQL queries

Established working login

Created a password authentication function for the login page that encrypts and decrypts user passwords making them undetectable to us as admins and any potential threats on security

--List frontend unit tests


--List backend unit tests


--Show updated documentation for your backend API 

- Updated main.go to include rest API to use HTTP requests. listen and serve on local host
    - Created reference to gin router that uses Get request
    - Ensures endpoints listening and waiting for requests
- Handler package contains individual endpoints to maintain organization
- Create package called newsfeed with functions that will return struct of type "Item" and will allow the ability to add items to a slice
- Added backend tests that test the Add() and GetAll() functions
    - Determines if newsfeed successfully adds an item to the slice of items
        - Checks length of slice
    - Determines if newsfeed successfully obtains all items within the slice
        - If the length of results slice do not equal 1
- Create newsfeed post and get request files
- Define interfaces within main.go that will return a slice of items and get all items
    - Will be used to integrate database
    - Used as function parameters
- References for creating REST API
    - https://youtu.be/LOn1GUsjOF4

Sprint 3 Backend Documentation:
- Implemented fiber framework to run server
    - Connected to database using MySQL Workbench and GORM
- Created and organized packages into folders
    - Includes controllers, databases, models, and routes
- Implemented Postman to send requests
- Created users
    - Register users function
    - Users struct to store each user's information
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
