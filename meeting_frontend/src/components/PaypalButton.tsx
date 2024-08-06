import { PayPalButtons, PayPalScriptProvider } from '@paypal/react-paypal-js';

const PayPalButton = () => {
  return (
    <PayPalScriptProvider options={{ "client-id": "your-paypal-client-id" }}>
      <PayPalButtons
        createOrder={(data, actions) => {
          return actions.order.create({
            purchase_units: [{
              amount: {
                value: '1.00',
              },
            }],
          });
        }}
        onApprove={(data, actions) => {
          return actions.order.capture().then((details) => {
            alert('Transaction completed by ' + details.payer.name.given_name);
          });
        }}
      />
    </PayPalScriptProvider>
  );
};

export default PayPalButton;
