function loadSales() {
    fetch('/sales')  
        .then(response => response.json())
        .then(data => {
            const salesTable = document.getElementById('salesTable'); 
            salesTable.innerHTML = ''; 

            data.forEach(sale => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${sale.id}</td>
                    <td>${sale.fecha}</td>
                    <td>${sale.total}</td>
                    <td><button onclick="editSale(${sale.id})">Edit</button></td>
                    <td><button onclick="deleteSale(${sale.id})">Delete</button></td>
                `;
                salesTable.appendChild(row);
            });
        })
        .catch(error => console.error('Error loading sales:', error));
}

function createSale(event) {
    event.preventDefault(); 

    const saleData = {
        venta_id: document.getElementById('ventaId').value,
        fecha: document.getElementById('fecha').value,
        total: document.getElementById('total').value
    };

    fetch('/sales', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(saleData) 
    })
        .then(response => response.json())
        .then(data => {
            alert('Sale created successfully');
            loadSales();  
        })
        .catch(error => console.error('Error creating sale:', error));
}

function editSale(saleId) {
    fetch(`/sales/${saleId}`)
        .then(response => response.json())
        .then(data => {
            document.getElementById('ventaId').value = data.venta_id;
            document.getElementById('fecha').value = data.fecha;
            document.getElementById('total').value = data.total;

            const submitButton = document.getElementById('submitSale');
            submitButton.textContent = 'Update Sale';
            submitButton.onclick = function(event) {
                updateSale(event, saleId);
            };
        })
        .catch(error => console.error('Error loading sale details:', error));
}

function updateSale(event, saleId) {
    event.preventDefault();

    const updatedSaleData = {
        venta_id: document.getElementById('ventaId').value,
        fecha: document.getElementById('fecha').value,
        total: document.getElementById('total').value
    };

    fetch(`/sales/${saleId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(updatedSaleData)
    })
        .then(response => response.json())
        .then(data => {
            alert('Sale updated successfully');
            loadSales();  
            const submitButton = document.getElementById('submitSale');
            submitButton.textContent = 'Create Sale';
            submitButton.onclick = createSale;
        })
        .catch(error => console.error('Error updating sale:', error));
}

function deleteSale(saleId) {
    if (confirm('Are you sure you want to delete this sale?')) {
        fetch(`/sales/${saleId}`, {
            method: 'DELETE'
        })
            .then(response => response.json())
            .then(data => {
                alert('Sale deleted successfully');
                loadSales();  
            })
            .catch(error => console.error('Error deleting sale:', error));
    }
}

document.addEventListener('DOMContentLoaded', function() {
    loadSales(); 

    const saleForm = document.getElementById('saleForm');
    saleForm.addEventListener('submit', createSale);
});
