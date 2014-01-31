import QtQuick 2.0

Item {
    id: window
    width: 600
    height: 400

    Component {
        id: imageDelegate
        Rectangle {
            width: window.width
            height: 25 
            color: ((index % 2 == 0) ? "#808080": "#999999")
            Row {
                spacing: 4
                anchors.left: parent.left
                anchors.leftMargin: 4
                children: [
                     Text {
                        text: imageModel.data(index).iD
                        elide: Text.ElideRight
			color: "red"
                     },
                     //Text {
                     //   text: imageModel.data(index).repoTags
                     //   elide: Text.ElideRight
                     //},
                     Text {
                        text: imageModel.data(index).created
                        elide: Text.ElideRight
                        color: "blue"
                     },
                     Text {
                        text: imageModel.data(index).size
                        elide: Text.ElideRight
                        color: "green"
                     },
                     Text {
                        text: imageModel.data(index).virtualSize
                        elide: Text.ElideRight
                        color: "white"
                     },
                     Text {
                        text: imageModel.data(index).parentId
                        elide: Text.ElideRight
                        color: "orange"
                     },
                     Text {
                        text: imageModel.data(index).repository
                        elide: Text.ElideRight
                        color: "red"
                     },
                     Text {
                        text: imageModel.data(index).tag
                        elide: Text.ElideRight
                     }
                ]
            }
        }
    }

    ListView {
        width: window.width
        model: imageModel.len
        delegate: imageDelegate 
        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.horizontalCenter: parent.horizontalCenter
    }
}
