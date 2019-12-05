require 'set'

input = <<-INPUT
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
INPUT
input = File.read('input.txt')
steps = input.split("\n")

class Node
  attr_accessor :value, :parents, :children

  def initialize(value)
    @value = value

    @parents = Set.new
    @children = Set.new
  end

  def add_parent(parent)
    parent.children << self
    parents << parent
  end

  def add_child(child)
    child.add_parent(self)
  end
end


STEP_RE = /Step (?<parent>[A-Z]) must be finished before step (?<child>[A-Z]) can begin\./
node_hash = {}
steps.each do |step|
  match = STEP_RE.match(step)
  parent_value = match['parent']
  parent_node = node_hash[parent_value] || Node.new(parent_value)
  node_hash[parent_value] = parent_node

  child_value = match['child']
  child_node = node_hash[child_value] || Node.new(child_value)
  node_hash[child_value] = child_node

  child_node.add_parent(parent_node)
end


def node_ready?(node, done_nodes)
  node.parents.all? { |p| done_nodes.include?(p) }
end

def all_ready?(nodes, done_nodes)
  nodes.all? { |node| node_ready?(node, done_nodes) }
end

def ready_nodes(nodes, done_nodes)
  nodes.select { |node| node_ready?(node, done_nodes) }
end

def all_done?(nodes, done_nodes)
  nodes.all? { |node| done_nodes.include?(node) }
end

step_list = ''
nodes = node_hash.values
done_nodes = Set.new
until all_done?(nodes, done_nodes)
  node = (ready_nodes(nodes, done_nodes) - done_nodes.to_a).min_by(&:value)
  step_list << node.value
  done_nodes << node
end

puts step_list
