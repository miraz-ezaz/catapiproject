# Cat API Project

This project is a Go-based web application built using the Beego framework. It interacts with The Cat API to display cat images, manage favorites, and allow users to vote on cats.

## Prerequisites

Before getting started, ensure that you have the following installed:

- Go (version 1.16 or later)
- Beego framework
- Bee CLI tool

## Getting Started

### 1. Clone the Repository

First, navigate to your `$GOPATH/src` directory, then clone the repository:

```bash
cd $GOPATH/src
git clone https://github.com/miraz-ezaz/catapiproject.git catapiproject
cd catapiproject
```

### 2. Install Beego and Bee CLI

To use Beego and the Bee CLI tool, you'll need to install them using the following commands:

```bash
go get github.com/beego/beego/v2
go install github.com/beego/bee/v2@latest
```

### 3. Install Project Dependencies

Next, you'll need to install the project dependencies. Run the following command:

```bash
go get ./...
```

After that, you should also tidy up your `go.mod` file:

```bash
go mod tidy
```

### 4. Update Configuration

Before running the project, you need to update the Beego configuration file to include your API key for The Cat API.

1. Open `conf/app.conf`.
2. Add your Cat API key and other settings:

```ini
appname = catapiproject
httpport = port_to_run_the_server
runmode = dev
catapiproject.apikey = your_cat_api_key
```

Replace `port_to_run_the_server` with the port number you want the server to run on (e.g., `8080`), and replace `your_cat_api_key` with your actual API key from The Cat API.

### 5. Running the Application

To run the application, use the Bee CLI tool:

```bash
bee run
```

This will start the Beego server, and you can access the application in your browser at `http://localhost:8080`.

## Project Structure

- `controllers/`: Contains all the controllers for handling web requests.
- `views/`: Contains the HTML templates.
- `static/`: Contains static files like CSS, JavaScript, and images.
- `conf/`: Contains the configuration files.
- `routers/`: Contains the routing configuration.

## Functionality

### 1. Voting Page

- **URL**: `/`
- **Functionality**: Displays a random cat image with options to upvote, downvote, or favorite the image. 
- **API Integration**: Uses The Cat API to fetch random cat images and handle user interactions.

### 2. Breeds Page

- **URL**: `/breeds`
- **Functionality**: Allows users to browse different cat breeds. Users can select a breed to view detailed information and a slideshow of images for that breed.
- **API Integration**: Fetches a list of breeds and breed-specific images from The Cat API.

### 3. Favorites Page

- **URL**: `/favorites`
- **Functionality**: Displays a gallery of the user's favorite cat images. Users can switch between grid and list views.
- **API Integration**: Fetches the user's favorite cat images from The Cat API and displays them in a scrollable gallery.

 
