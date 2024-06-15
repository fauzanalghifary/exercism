class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    items.map { |item| item[:name] }.sort
  end

  def cheap
    cheap_limit = 30
    items.select { |item| item[:price] < cheap_limit }
  end

  def out_of_stock
    items.select { |item| item[:quantity_by_size].empty? }
  end

  def stock_for_item(name)
    items.find { |item| item[:name] == name }[:quantity_by_size]
  end

  def total_stock
    # items.map { |item| item[:quantity_by_size].map { |_, num| num } }.flatten.sum
    @items.sum { |item| item[:quantity_by_size].values.sum }
  end

  private
  attr_reader :items
end
