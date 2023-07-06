
  async function registrarControl() {
   let datos = {};
     datos.descripcion = document.getElementById('txtDescripcion').value;

          const request = await fetch('control', {
              method: 'POST',
              body: JSON.stringify(datos),
              headers: {
                  'Accept': 'application/json',
                  'Content-Type': 'application/json'
              },
          })
        if (request.status == 200) {
            alert("Control registrado correctamente");
        }
  }
