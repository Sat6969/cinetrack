const API_URL = import.meta.env.VITE_API_URL;

async function registerUser(name, email, password) {
  const response = await fetch(
    API_URL + "/register",
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json",
      },

      body: JSON.stringify({
        name: name,
        email: email,
        password: password,
      }),
    }
  );

  const data = await response.json();

  if (!response.ok) {
    throw new Error(
      data.error || "registration failed"
    );
  }

  return data;
}

async function loginUser(email, password) {
  const response = await fetch(
    API_URL + "/login",
    {
      method: "POST",

      headers: {
        "Content-Type": "application/json",
      },

      body: JSON.stringify({
        email: email,
        password: password,
      }),
    }
  );

  const data = await response.json();

  if (!response.ok) {
    throw new Error(
      data.error || "login failed"
    );
  }

  localStorage.setItem("token", data.token);

  return data;
}

function getToken() {
  return localStorage.getItem("token");
}

function logoutUser() {
  localStorage.removeItem("token");
}

export {
  registerUser,
  loginUser,
  getToken,
  logoutUser,
};