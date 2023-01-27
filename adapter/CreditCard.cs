namespace  AdapterPattern{
    public class CreditCard : ICard
    {
        public string GetCardDetail(int customerNumber)
        {
            return $"Credit Card Detail for {customerNumber}";
        }

        public int GetTotalAmount(int cardNumber)
        {
            return 1000;
        }
    }
}