class Node
  attr_reader :value, :left, :right

  def initialize(value, left, right)
    @value = value
    @left = left
    @right = right
  end

  def ==(other)
    return false unless other.is_a?(Node)
    value == other.value && left == other.left && right == other.right
  end
end

class Zipper
  attr_reader :focus, :breadcrumbs

  def self.from_tree(tree)
    new(tree, [])
  end

  def initialize(focus, breadcrumbs)
    @focus = focus
    @breadcrumbs = breadcrumbs
  end

  def value
    @focus.value
  end

  def left
    return nil if @focus.left.nil?
    new_breadcrumb = { direction: :left, parent_value: @focus.value, sibling: @focus.right }
    Zipper.new(@focus.left, @breadcrumbs + [new_breadcrumb])
  end

  def right
    return nil if @focus.right.nil?
    new_breadcrumb = { direction: :right, parent_value: @focus.value, sibling: @focus.left }
    Zipper.new(@focus.right, @breadcrumbs + [new_breadcrumb])
  end

  def up
    return nil if @breadcrumbs.empty?

    crumb = @breadcrumbs.last
    parent_node = if crumb[:direction] == :left
                    Node.new(crumb[:parent_value], @focus, crumb[:sibling])
                  else
                    Node.new(crumb[:parent_value], crumb[:sibling], @focus)
                  end

    Zipper.new(parent_node, @breadcrumbs[0...-1])
  end

  def set_value(new_value)
    new_focus = Node.new(new_value, @focus.left, @focus.right)
    Zipper.new(new_focus, @breadcrumbs)
  end

  def set_left(new_left)
    new_focus = Node.new(@focus.value, new_left, @focus.right)
    Zipper.new(new_focus, @breadcrumbs)
  end

  def set_right(new_right)
    new_focus = Node.new(@focus.value, @focus.left, new_right)
    Zipper.new(new_focus, @breadcrumbs)
  end

  def to_tree
    current = self
    current = current.up while current.breadcrumbs.any?
    current.focus
  end

  def ==(other)
    return false unless other.is_a?(Zipper)
    focus == other.focus && breadcrumbs == other.breadcrumbs
  end
end
