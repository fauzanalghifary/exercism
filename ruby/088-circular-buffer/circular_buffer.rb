class CircularBuffer
  def initialize(capacity)
    @capacity = capacity
    @buffer = Array.new(capacity)
    @read_pos = 0
    @write_pos = 0
    @size = 0
  end

  def read
    raise BufferEmptyException if empty?
    item = @buffer[@read_pos]
    @buffer[@read_pos] = nil
    @read_pos = (@read_pos + 1) % @capacity
    @size -= 1
    item
  end

  def write(item)
    raise BufferFullException if full?
    write_item(item)
  end

  def write!(item)
    read if full?
    write_item(item)
  end

  def clear
    @buffer = Array.new(@capacity)
    @read_pos = 0
    @write_pos = 0
    @size = 0
  end

  private

  def write_item(item)
    return if item.nil?
    @buffer[@write_pos] = item
    @write_pos = (@write_pos + 1) % @capacity
    @size += 1
  end

  def full?
    @size == @capacity
  end

  def empty?
    @size.zero?
  end

  class BufferEmptyException < StandardError; end
  class BufferFullException < StandardError; end
end