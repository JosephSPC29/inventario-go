const apiBaseUrl = "http://localhost:8080";

async function fetchUsers() {
  try {
    const response = await fetch(`${apiBaseUrl}/users`);
    const users = await response.json();

    const tableBody = document.querySelector("#usersTable tbody");
    tableBody.innerHTML = ""; 

    users.forEach(user => {
      const row = document.createElement("tr");
      row.innerHTML = `
        <td>${user.id}</td>
        <td>${user.username}</td>
        <td>${user.email}</td>
        <td>${user.role}</td>
      `;
      tableBody.appendChild(row);
    });
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}

async function createUser(event) {
  event.preventDefault();

  const form = document.querySelector("#userForm");
  const formData = new FormData(form);

  const user = {
    username: formData.get("username"),
    email: formData.get("email"),
    role: formData.get("role"),
  };

  try {
    const response = await fetch(`${apiBaseUrl}/users`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });

    if (response.ok) {
      alert("Usuario registrado con éxito");
      form.reset();
      fetchUsers(); 
    } else {
      const error = await response.json();
      alert(`Error: ${error.message}`);
    }
  } catch (error) {
    console.error("Error creating user:", error);
  }
}

document.addEventListener("DOMContentLoaded", () => {
  fetchUsers();

  const form = document.querySelector("#userForm");
  form.addEventListener("submit", createUser);
});
