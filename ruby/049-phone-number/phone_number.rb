class PhoneNumber

  # @param [string] phone_string
  def self.clean(phone_string)
    phone_num = phone_string.gsub(/^1|^\+1/, '').gsub(/[^0-9]/, '')
    return nil if phone_num.size != 10 || phone_num.to_s[0].to_i < 2 || phone_num.to_s[3].to_i < 2

    phone_num
  end


  # VALID_PHONE_NUMBER = /^([2-9]\d\d){2}\d{4}$/.freeze
  # def self.clean(phone_string)
  #   phone_num = phone_string.gsub(/\D/, '').sub(/^1/, '')
  #   phone_num[VALID_PHONE_NUMBER]
  # end
end