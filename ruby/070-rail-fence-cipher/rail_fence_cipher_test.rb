require 'minitest/autorun'
require_relative 'rail_fence_cipher'

class RailFenceCipherTest < Minitest::Test
  def test_encode_with_empty_string
    assert_equal '', RailFenceCipher.encode('', 4)
  end

  def test_encode_with_one_rail
    assert_equal 'One rail, only one rail',
      RailFenceCipher.encode('One rail, only one rail', 1)
  end

  def test_encode_with_two_rails
    assert_equal 'XXXXXXXXXOOOOOOOOO',
      RailFenceCipher.encode('XOXOXOXOXOXOXOXOXO', 2)
  end

  def test_encode_with_three_rails
    assert_equal 'WECRLTEERDSOEEFEAOCAIVDEN',
      RailFenceCipher.encode('WEAREDISCOVEREDFLEEATONCE', 3)
  end

  def test_encode_with_ending_in_the_middle
    assert_equal 'ESXIEECSR', RailFenceCipher.encode('EXERCISES', 4)
  end

  def test_encode_with_less_letters_than_rails
    assert_equal 'More rails than letters',
      RailFenceCipher.encode('More rails than letters', 24)
  end

  def test_decode_with_empty_string
    assert_equal '', RailFenceCipher.decode('', 4)
  end

  def test_decode_with_one_rail
    assert_equal 'ABCDEFGHIJKLMNOP',
      RailFenceCipher.decode('ABCDEFGHIJKLMNOP', 1)
  end

  def test_decode_with_two_rails
    assert_equal 'XOXOXOXOXOXOXOXOXO',
      RailFenceCipher.decode(+'XXXXXXXXXOOOOOOOOO', 2)
  end

  def test_decode_with_three_rails
    assert_equal 'THEDEVILISINTHEDETAILS',
      RailFenceCipher.decode(+'TEITELHDVLSNHDTISEIIEA', 3)
  end
end
