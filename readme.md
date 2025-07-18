# Welcome to this Blog Platform!


# Setup instructions
In order to make this blog platform, I first created the architecture of the platform by creating all of folders to organise the code. I initialised a go module to manage all of my different packages. I initialised the database in the db folder and called on .env to use my secure data. I then created my middleware to generate and verify the JWT, and ensure that the user is authenticated. I also included a function to extract the user ID and user role from the JWT which will be checked against the information inputted to ensure that the user is authorised for certain features of the blog, and also to ensure that they are only about to edit or delete their own posts. In the config file, I created a function using a switch statement to determine what access the user has based on their role. In my utils folder I created functions to hash the password of the registered user and a compare function to check that the passowrd that has been inputted is correct against the database. I then created my user and post structs with all the relevant fields that I wanted to include. After this, I created all of the features I wanted the blog platform to have: register user, login, create post, get all posts, get a specific post, update own post, and delete own post. I created these features by using the handlers, services and repository layers and then added all the routes in the router folder. Finally, I initialised all of these in the main.go and then started the server. 

# Technologies used
I used Golang to write the code, Github to host my git repository and Render to deploy my blog platform.


# Deployment instructions
To use the blog API, please first register as a user so that you can start creating your first blog post! 
To do this, enter your name, email and password and use the /register route.
Then you will need to log in with the same details by using the /login route. 
After this, you can start creating your own posts, using /posts and inputting the title and content of your blog post.
If you want to see all of the blog posts that have been uploaded, you can do so by using the same /posts route. 
You can also get a specific post, update your own post, or delete your own post by using /posts/: and enter the id of the post you want. 