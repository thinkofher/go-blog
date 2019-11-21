FONT_VERSION="1.1.1"
DOWNLOAD_LINK="https://github.com/iconic/open-iconic/archive/${FONT_VERSION}.tar.gz"
FONT_LOCALIZATION="static"

wget $DOWNLOAD_LINK
tar xzf "${FONT_VERSION}.tar.gz"
mkdir -p $FONT_LOCALIZATION
cp -ru "open-iconic-${FONT_VERSION}/font" $FONT_LOCALIZATION/

rm -rf "open-iconic-${FONT_VERSION}"
rm -rf $FONT_VERSION.tar.gz*
