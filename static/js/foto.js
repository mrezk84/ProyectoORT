document.addEventListener("DOMContentLoaded", function () {
    const baseURL = "http://localhost:5000/fotos/formulario"; 
    const photoTable = document.getElementById("fotosTBody");

    // Obtener la lista de fotos desde la API
    fetch("/api/fotos", { method: "GET" })
        .then(response => response.json())
        .then(data => {
            // Recorrer la lista de fotos y agregarlas a la tabla
            data.forEach(photo => {
                const row = document.createElement("tr");

                // ID
                const idCell = document.createElement("td");
                idCell.textContent = photo.ID;
                row.appendChild(idCell);

                // Nombre
                const nombreCell = document.createElement("td");
                nombreCell.textContent = photo.Nombre;
                row.appendChild(nombreCell);

                // Notas
                const notasCell = document.createElement("td");
                notasCell.textContent = photo.Notas;
                row.appendChild(notasCell);

                // Formulario ID
                const formularioIdCell = document.createElement("td");
                formularioIdCell.textContent = photo.FormularioID;
                row.appendChild(formularioIdCell);

                // Imagen
                 // Imagen
                 const imagenCell = document.createElement("td");
                 const imagen = document.createElement("img");
                 const imageURL = baseURL + "/fotos/" + photo.Nombre; // Reemplaza con el campo que contiene el nombre del archivo de imagen
                 imagen.src = imageURL;
                 imagen.alt = "Imagen"; // Puedes ajustar el texto alternativo según tus necesidades
                 imagenCell.appendChild(imagen);
                 row.appendChild(imagenCell);
 
                 photoTable.appendChild(row);
            });
        })
        .catch(error => {
            console.error("Error al obtener las fotos:", error);
        });
});
