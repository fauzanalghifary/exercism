class Node
  attr_accessor :next, :prev
  attr_reader :value

  # @param[Numeric]value
  def initialize(value)
    @value = value
    @next = nil
    @prev = nil
  end
end

class Deque
  def initialize
    @head = nil
    @tail = nil
    @size = 0
  end

  def push(num)
    new_node = Node.new(num)

    if @size == 0
      @head = @tail = new_node
      @size += 1
      return
    end

    @size += 1
    prev_tail = @tail
    prev_tail.next = new_node
    new_node.prev = @tail
    @tail = new_node
  end

  def unshift(num)
    new_node = Node.new(num)

    if @size == 0
      @head = @tail = new_node
      @size += 1
      return
    end

    @size += 1
    prev_head = @head
    prev_head.prev = new_node
    new_node.next = @head
    @head = new_node
  end

  def pop
    return ArgumentError if @size == 0
    val = @tail.value

    if @size == 1
      @head = nil
      @tail = nil
      @size -= 1
      return val
    end

    @size -= 1
    next_tail = @tail.prev
    next_tail.next = nil
    @tail = next_tail

    val
  end

  def shift
    return ArgumentError if @size == 0
    val = @head.value

    if @size == 1
      @head = nil
      @tail = nil
      @size -= 1
      return val
    end

    @size -= 1
    next_head = @head.next
    next_head.prev = nil
    @head = next_head

    val
  end
end