import QtQuick 2.0

Item {
    width: 320; height: 200

    ListView {
        width: 120;
        model: images.len
        delegate: Text {
            text: images.image(index)
        }
        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.horizontalCenter: parent.horizontalCenter
    }
}
