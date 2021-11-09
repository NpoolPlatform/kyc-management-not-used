FROM jrei/systemd-centos:7

COPY KycManagement.viper.yaml /etc/KycManagement/KycManagement.viper.yaml
COPY kyc-management /usr/local/bin/kyc-management

CMD ["kyc-management", "run"]
