document.addEventListener("DOMContentLoaded", function() {
    const fetchButton = document.getElementById("fetchButton");
    const orderIdInput = document.getElementById("orderId");
    const orderDetails = document.getElementById("orderDetails");
    const orderData = document.getElementById("orderData");

    fetchButton.addEventListener("click", async function() {
        const orderId = orderIdInput.value;
        if (!orderId) {
            return;
        }

        const response = await fetch(`/get_order/${orderId}`);
        if (response.status === 302) {
            const order = await response.json();
            orderData.textContent = JSON.stringify(order, null, 2);
            orderDetails.classList.remove("hidden");
        } else {
            orderDetails.classList.add("hidden");
            alert("Order not found");
        }
    });
});
