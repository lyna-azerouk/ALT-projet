
const loadMenus = () => {
    const data = require("./data.json");
    return data.menus.map((menu) => {
        // list to string seperated by :
        const apiMenu = {}
        apiMenu["name"] = menu["name"]
        apiMenu["price"] = menu["price"]
        apiMenu["description"] = menu["ingredients"].join(":")
        apiMenu["image"] = menu["image_url"]
        return apiMenu;
    })
}

const token0riginel = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTU4NjAzNzEyLCJleHAiOjE3MTE3NDA4OTMsImlhdCI6MTcxMTY1NDQ5M30.5T7T7_wb17i2JbVrANoqUkZRuh-elwT7oD_ZaSSnIHg"

const getRandomMenus = (menus, n) => {
    const randomMenus = [];
    // do no duplicate
    const indexes = new Set();
    while (indexes.size < n) {
        indexes.add(Math.floor(Math.random() * menus.length));
    }
    indexes.forEach((index) => randomMenus.push(menus[index]));
    return randomMenus;
}

const main = () => {
    const myHeaders = new Headers();
    myHeaders.append("Authorization", "Bearer " + token0riginel);

    const requestOptions = {
        method: "GET",
        headers: myHeaders,
        redirect: "follow"
    };

    const menus = loadMenus();

    fetch("http://localhost:8080/restaurants/2.3545681577381847/48.84621136872197/0.001", requestOptions)
        .then((response) => response.json())
        .then((result) => {
            result["restaurants"].forEach(async (restaurant) => {
                const token = await registerRestaurant(restaurant);
                const randomMenus = getRandomMenus(menus, 5);
                if (token !== "") {
                    randomMenus.forEach((menu) => {
                        registerMenu(restaurant["id"], menu, token);
                    });
                }
            });
        })
        .catch((error) => console.error(error));
}


const registerMenu = (restaurantId, menu, token) => {
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Authorization", "Bearer " + token);

    const raw = JSON.stringify(menu);

    const requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow"
    };

    fetch("http://localhost:8080/restaurant/menu/" + restaurantId, requestOptions)
        .then((response) => response.text())
        .catch((error) => console.error(error));
}


/** Register a restaurant in the database
 * 
 * @param {*} restaurant 
 * @returns 
 */
const registerRestaurant = async (restaurant) => {
    const id = restaurant["id"];
    const psw = String(id)
    const name = restaurant["tags"]["name"];


    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    const raw = JSON.stringify({
        "id": id,
        "password": psw,
        "name": name
    });

    const requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow"
    };

    var token = "";

    fetch("http://localhost:8080/signup/restaurant", requestOptions)
        .then((response) => response.text())
        .catch((error) => { console.error(error) });


    // login
    const credentials = {
        "id": id,
        "password": psw
    }

    const myHeaders2 = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const raw2 = JSON.stringify(credentials);

    const requestOptions2 = {
        method: "POST",
        headers: myHeaders2,
        body: raw2,
        redirect: "follow"
    };

    const request = await fetch("http://localhost:8080/auth/restaurant", requestOptions2);
    const response = await request.json();
    token = await response["token"];

    return token;

}


const playground = () => {
    // load data.json file
    console.log(getRandomMenus(loadMenus(), 5));

}

main();