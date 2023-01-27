namespace AdapterPattern{
    public class PrePaidCard
    {
        public string GetPrepaidDetail(string customerNumber)
        {
            return $"PrePaid Card Detail for {customerNumber}";
        }

        public int GetAmount(string cardNumber)
        {
            return 10000;
        }
    }
}