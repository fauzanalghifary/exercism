class BankAccount
  def initialize
    @balance = 0
    @open = false
  end

  def open
    raise ArgumentError, "You can't open an already open account" if @open
    @open = true
    @balance = 0
  end

  def close
    raise ArgumentError, "You can't close an already closed account" unless @open
    @open = false
  end

  def balance
    raise ArgumentError, "You can't check the balance of a closed account" unless @open
    @balance
  end

  def deposit(amount)
    raise ArgumentError, "You can't deposit money into a closed account" unless @open
    raise ArgumentError, "You can't deposit a negative amount" if amount < 0
    @balance += amount
  end

  def withdraw(amount)
    raise ArgumentError, "You can't withdraw money into a closed account" unless @open
    raise ArgumentError, "You can't withdraw a negative amount" if amount < 0
    raise ArgumentError, "You can't withdraw more than you have" if amount > @balance
    @balance -= amount
  end
end
