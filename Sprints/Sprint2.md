Detail work you've completed in Sprint 2
---
- Frontend: Implemented full functionality of the login, sign up, and homepage. Generated a logo and linked it as a button to the homepage. Added skeletons for future implementation of admin and user roles with seperate functionality and access for each.

- Backend: Created REST api and integrated front end with backend. Created unit tests. 

List unit tests and Cypress test for frontend
---
- Cypress Test's - 

1. Scans and clicks the sign up button in the top left of the homepage and verifies that the user has visited the register page

2. Scans and clicks the log in button in the top left of the homepage and verifies that the user has visted the log in page


List unit tests for backend
---
- Our unit tests check to see if there has been an item added to the feed.


Add documentation for your backend API
---

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

---