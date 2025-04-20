class Element
  attr_accessor :next
  attr_reader :value

  # @param[Numeric]value
  def initialize(value)
    @value = value
    @next = nil
  end

  def datum
    @value
  end
end

class SimpleLinkedList
  attr_reader :head

  def initialize(arr = [])
    @head = nil
    arr.each { |value| internal_push_value(value) }
  end

  # @param[Element] element
  def push(element)
    element.next = @head
    @head = element
    self
  end

  # @return[Element, nil]
  def pop
    return nil unless @head

    popped_element = @head
    @head = @head.next
    popped_element.next = nil
    popped_element
  end

  # @return[Array]
  def to_a
    result = []
    current = @head
    while current
      result << current.datum
      current = current.next
    end
    result
  end

  # @return[SimpleLinkedList]
  def reverse!
    previous = nil
    current = @head
    while current
      next_node = current.next  # Store next node
      current.next = previous   # Reverse current node's pointer
      previous = current        # Move previous one step forward
      current = next_node       # Move current one step forward
    end
    @head = previous            # The last 'previous' is the new head
    self
  end

  private

  def internal_push_value(value)
    element = Element.new(value)
    push(element)
  end
end
