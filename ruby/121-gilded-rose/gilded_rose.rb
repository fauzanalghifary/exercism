Item = Struct.new(:name, :sell_in, :quality)

class GildedRose
  def initialize(items)
    @items = items
  end

  def update!
    @items.each do |item|
      update_item(item)
    end
  end

  private

  def conjured?(item)
    item.name.start_with?("Conjured")
  end

  def sulfuras?(item)
    item.name.include?("Sulfuras, Hand of Ragnaros")
  end

  def aged_brie?(item)
    item.name.include?("Aged Brie")
  end

  def backstage_pass?(item)
    item.name.include?("Backstage passes to a TAFKAL80ETC concert") ||
      item.name.include?("backstage passes to a TAFKAL80ETC concert")
  end

  def update_item(item)
    # Regular Sulfuras never changes
    return if sulfuras?(item) && !conjured?(item)

    # Store original sell_in for quality calculations
    original_sell_in = item.sell_in

    # Update quality based on item type BEFORE decrementing sell_in
    if aged_brie?(item) && !conjured?(item)
      update_aged_brie_quality(item, original_sell_in)
    elsif backstage_pass?(item) && !conjured?(item)
      update_backstage_pass_quality(item, original_sell_in)
    elsif backstage_pass?(item) && conjured?(item)
      update_conjured_backstage_pass_quality(item, original_sell_in)
    elsif conjured?(item)
      update_conjured_quality(item, original_sell_in)
    else
      update_normal_quality(item, original_sell_in)
    end

    # Decrement sell_in (Conjured Sulfuras has a sell-by date)
    item.sell_in -= 1
  end

  def update_aged_brie_quality(item, original_sell_in)
    if original_sell_in > 0
      increase_quality(item, 1)
    else
      increase_quality(item, 2)
    end
  end

  def update_backstage_pass_quality(item, original_sell_in)
    if original_sell_in <= 0
      item.quality = 0
    elsif original_sell_in <= 5
      increase_quality(item, 3)
    elsif original_sell_in <= 10
      increase_quality(item, 2)
    else
      increase_quality(item, 1)
    end
  end

  def update_conjured_backstage_pass_quality(item, original_sell_in)
    # Conjured items drop to 0 once sell-by date has arrived
    if original_sell_in <= 0
      item.quality = 0
      return
    end

    # Conjured backstage passes increase by 1 less than non-conjured
    if original_sell_in <= 5
      increase_quality(item, 2)
    elsif original_sell_in <= 10
      increase_quality(item, 1)
    end
    # > 10 days: no change (1 less than the normal +1)
  end

  def update_conjured_quality(item, original_sell_in)
    # Conjured items drop to 0 once sell-by date has arrived
    if original_sell_in <= 0
      item.quality = 0
      return
    end

    # Conjured normal items (including Aged Brie and Sulfuras) degrade twice as fast
    if aged_brie?(item)
      increase_quality(item, 1)
    elsif sulfuras?(item)
      # Conjured Sulfuras doesn't change before sell date
      return
    else
      decrease_quality(item, 2)
    end
  end

  def update_normal_quality(item, original_sell_in)
    if original_sell_in > 0
      decrease_quality(item, 1)
    else
      decrease_quality(item, 2)
    end
  end

  def increase_quality(item, amount)
    item.quality = [item.quality + amount, 50].min
  end

  def decrease_quality(item, amount)
    item.quality = [item.quality - amount, 0].max
  end
end
