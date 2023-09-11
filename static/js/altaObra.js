
  async function registrarObra() {
   let datos = {};
     datos.nombre = document.getElementById('txtNombre').value;
          const request = await fetch("http://localhost:5000/obras/registrar", {
              method: 'POST',
              body: JSON.stringify(datos),
              headers: {
                  'Accept': 'application/json',
                  'Content-Type': 'application/json'
              },
          })
        if (request.status == 200) {
            alert("Obra registrada correctamente");
        }else{
            alert("Error registrando la obra");
        }
            window.location.href = `obras.html`;
  }
