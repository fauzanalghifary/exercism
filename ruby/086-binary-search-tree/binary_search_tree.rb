class Node
  attr_accessor :value, :right, :left
  def initialize(value)
    @value = value
    @right = @left = nil
  end

  def data
    @value
  end
end

class Bst
  include Enumerable

  def initialize(value)
    @root = Node.new(value)
    @root
  end

  def insert(value, current_node = @root)
    if current_node.nil?
      return Node.new(value)
    end

    if value <= current_node.value
      current_node.left = insert(value, current_node.left)
    else
      current_node.right = insert(value, current_node.right)
    end

    current_node
  end

  def each(&block)
    # Standard pattern: return Enumerator if no block is given.
    return enum_for(:each) unless block_given?

    in_order_traversal(@root, &block)
  end

  def data
    @root.data
  end

  def left
    @root.left
  end

  def right
    @root.right
  end

  private

  def in_order_traversal(node, &block)
    return unless node                      # Base case: stop if node is nil

    in_order_traversal(node.left, &block)   # Traverse left subtree
    block.call(node.data)                   # Yield current node's data
    in_order_traversal(node.right, &block)  # Traverse right subtree
  end
end